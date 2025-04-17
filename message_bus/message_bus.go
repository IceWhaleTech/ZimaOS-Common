package message_bus

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	socketio "github.com/maldikhan/go.socket.io/socket.io/v5/client"
	"github.com/rfyiamcool/backoff"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// 这里为什么不用socketio默认的重连，是因为如果error或者disconnect，对应的msg port可能会也发生变化，自动重连连不上
// 然后为什么不用gateway来固定url，可以看下面的注释，主要是因为这个sokcetio有点坑。
var needToReconnectMsg = sync.WaitGroup{}

type MessageBusService struct {
	client    *socketio.Client
	listeners map[string][]func(Event)
	rooms     []string
	backoff   *backoff.Backoff
	lock      sync.RWMutex
}

type loggerImpl struct{}

func (l loggerImpl) Debugf(format string, v ...any) {
}

func (l loggerImpl) Infof(format string, v ...any) {
}

func (l loggerImpl) Warnf(format string, v ...any) {
}

func (l loggerImpl) Errorf(format string, v ...any) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic in connect", zap.Any("error", r))
			needToReconnectMsg = sync.WaitGroup{}
		}
	}()

	needToReconnectMsg.Done()

	logger.Error(fmt.Sprintf(format, v...))
	if format == "receiveWsError: %v" && len(v) > 0 {
		errStr := ""
		switch e := v[0].(type) {
		case string:
			errStr = e
		case error:
			errStr = e.Error()
		}

		if errStr == "EOF" {
			logger.Info("msg bus got unexpected EOF, try to reconnect")
			needToReconnectMsg.Done()
		}
	}
}

func NewMessageBusService(
	rooms []string,
) *MessageBusService {

	b := backoff.NewBackOff(
		backoff.WithMinDelay(1*time.Second),
		backoff.WithMaxDelay(20*time.Second),
		backoff.WithFactor(2),
	)

	service := &MessageBusService{
		listeners: make(map[string][]func(Event)),
		backoff:   b,
		rooms:     rooms,
	}
	go func() {

		// 这里使用避火算法
		for {
			// 这里设成nil，是为了解决重连时msg可能的发送失败
			// 如果不是nil，但是这个client断了，就会直接返回错误
			// 如果是nil，就会等待发送，而不是直接失败
			service.lock.Lock()
			service.client = nil
			service.lock.Unlock()

			service.connect()
			time.Sleep(b.Duration())
		}
	}()
	return service
}

func (s *MessageBusService) connect() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic in connect", zap.Any("error", r))
			needToReconnectMsg = sync.WaitGroup{}
		}
	}()

	messageURLFile := "/var/run/casaos/message-bus.url"
	messageURL, err := os.ReadFile(messageURLFile)
	if err != nil {
		log.Printf("Failed to read message URL: %v", err)
	}

	// 这里不走gateway连接而是直接连接message是因为有以下原因
	// 1. gateway本身重启频率比msg bus高(注册路由机制)
	// 2. 发现msg变化之后，这个client使用socketio的重连速度比较慢。
	// 2.1 前端用的socketio，重连速度比较快。
	// 后续可以试试看这个 https://github.com/zishang520/socket.io-client-go
	// 这个是下面业务代码的时候还没有出，是在issue催了之后，刚刚推出的

	client, err := socketio.NewClient(
		socketio.WithRawURL(fmt.Sprintf("%s/v2/message_bus/socket.io", messageURL)),
		socketio.WithLogger(loggerImpl{}),
	)

	if err != nil {
		log.Printf("Failed to create socket.io client: %v", err)
		return
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Printf("Failed to connect to message bus: %v", err)
		return
	}

	client.On("connect", func() {
		for _, room := range s.rooms {
			err := s.JoinRoom(room)
			if err != nil {
				logger.Error("Failed to join room", zap.String("room", room), zap.Error(err))
			}
		}
	})

	needToReconnectMsg.Add(1)

	s.backoff.Reset()
	client.On("disconnect", func() {
		logger.Error("Message bus disconnected, try to reconnect")
		needToReconnectMsg.Done()
	})

	client.On("error", func() {
		logger.Error("Message bus error, try to reconnect")
		needToReconnectMsg.Done()
	})

	// 为什么需要给 AddEventHandler 传 client，是因为要先监听完所有已add的事件，才允许AddEventHandler新的事件，避免下面的case3
	// case1. message bus先连上，listeners为空，然后一个个添加addEventHandler。没有问题
	// case2. message bus没连上，然后一个个添加addEventHandler到s.listeners，等连上的时候重新监听所有的事件
	// case3. message bus连上了，为了避免 s.client = client 之后cpu切出去执行 AddEventHandler，然后再回来执行 for ，导致重复监听

	// 把 listeners 重新监听上
	for eventType, handlers := range s.listeners {
		for _, handler := range handlers {
			AddEventHandler(eventType, handler, client)
		}
	}

	s.lock.Lock()
	s.client = client
	s.lock.Unlock()
	logger.Info("Message bus connected")

	needToReconnectMsg.Wait()
}

func AddEventHandler(eventName string, handler func(Event), targetClient *socketio.Client) {
	logger.Info("add event handler", zap.Any("event name", eventName))

	targetClient.On(eventName, func(event interface{}) {
		rawEvent := event.(map[string]interface{})

		var msgEvent Event
		msgEvent.Name = eventName

		if _, ok := rawEvent["SourceID"]; ok {
			msgEvent.SourceID = rawEvent["SourceID"].(string)
		}

		if _, ok := rawEvent["Properties"]; ok {
			rawPropertise := rawEvent["Properties"].(map[string]interface{})
			// string json to map
			var properties map[string]string = make(map[string]string)
			for key, value := range rawPropertise {
				properties[key] = value.(string)
			}
			msgEvent.Properties = properties
		}

		if _, ok := rawEvent["uuid"]; ok {
			msgEvent.Uuid = lo.ToPtr(rawEvent["uuid"].(string))
		}

		if _, ok := rawEvent["Timestamp"]; ok {
			rawTimestamp := rawEvent["Timestamp"].(float64)
			timestamp, cerr := time.Parse(time.RFC3339, fmt.Sprintf("%f", rawTimestamp))
			if cerr == nil {
				msgEvent.Timestamp = lo.ToPtr(timestamp)
			}
		}
		handler(msgEvent)
	})
}

func (s *MessageBusService) AddEventHandler(eventName string, handler func(Event)) {
	s.listeners[eventName] = append(s.listeners[eventName], handler)

	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.client == nil {
		// 如果client没有建立，先加到listeners。等msg bus连上了再初始化

		// 上面是先把现在的handlers都加到client，才添加到client。
		// 避免这里添加一次事件，然后初始化又添加一次的双重监听
		logger.Info("client is nil, wait to add event handler", zap.Any("event name", eventName))
		return
	}

	AddEventHandler(eventName, handler, s.client)
}

func (s *MessageBusService) JoinRoom(room string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.client == nil {
		return fmt.Errorf("message bus not connected")
	}
	s.rooms = append(s.rooms, room)
	return s.client.Emit("room:join", room)
}

func (s *MessageBusService) LeaveRoom(room string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.client == nil {
		return fmt.Errorf("message bus not connected")
	}
	s.rooms = lo.Filter(s.rooms, func(r string, _ int) bool {
		return r != room
	})
	return s.client.Emit("room:leave", room)
}

func (s *MessageBusService) PublishEvent(eventType EventType, properties map[string]string) error {
	// 如果发消息的时候，msg 还没有连上，就先等待
	// 如果尝试超过3次仍然无法连接，则放弃
	attempts := 0
	maxAttempts := 3

	for {
		s.lock.RLock()
		if s.client != nil {
			break
		}
		s.lock.RUnlock()
		attempts++
		if attempts > maxAttempts {
			return fmt.Errorf("failed to connect to message bus, tried %d times", maxAttempts)
		}
		time.Sleep(2 * time.Second)
	}

	room := "event"
	if eventType.Room != "" {
		room = eventType.Room
	}

	// 如果报错是 connection reset by peer，那么说明 properties 太大了,超过buffer了
	// 后面这里应该限制事件的大小，比如不允许事件超过多少k。
	err := s.client.Emit(eventType.Name, eventType.SourceID, properties, room)
	s.lock.RUnlock()
	return err
}
