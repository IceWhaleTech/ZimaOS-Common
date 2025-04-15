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
)

var wg = sync.WaitGroup{}

type MessageBusService struct {
	client    *socketio.Client
	listeners map[string][]func(Event)
	backoff   *backoff.Backoff
}

type loggerImpl struct{}

func (l loggerImpl) Debugf(format string, v ...any) {
}

func (l loggerImpl) Infof(format string, v ...any) {
}

func (l loggerImpl) Warnf(format string, v ...any) {
}

func (l loggerImpl) Errorf(format string, v ...any) {
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
			logger.Info("try to reconnect")
			wg.Done()
		}
	}
}

func NewMessageBusService() *MessageBusService {
	b := backoff.NewBackOff(
		backoff.WithMinDelay(1*time.Second),
		backoff.WithMaxDelay(20*time.Second),
		backoff.WithFactor(2),
	)

	service := &MessageBusService{
		listeners: make(map[string][]func(Event)),
		backoff:   b,
	}
	go func() {
		// 这里使用避火算法
		for {
			service.client = nil
			service.connect()
			time.Sleep(b.Duration())
		}
	}()
	return service
}

func (s *MessageBusService) connect() {
	messageURLFile := "/var/run/casaos/message-bus.url"
	messageURL, err := os.ReadFile(messageURLFile)
	if err != nil {
		log.Printf("Failed to read message URL: %v", err)
	}

	// 不走gateway连接，遇到两个奇怪的问题
	// 1. 第一次连上之后会马上拿到EOF
	// 2. 关掉msg bus，通过gateway的连接还能连上，但是不会报错或者啥
	// 后续看看原因是什么

	// 这个socketio的自动重连没生效，不知道为什么

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

	wg.Add(1)

	s.backoff.Reset()
	client.On("disconnect", func() {
		logger.Error("Message bus disconnected,尝试重连")
		wg.Done()
	})

	client.On("error", func() {
		logger.Error("Message bus error")
		wg.Done()
	})

	logger.Info("Message bus connected")

	// 把listeners重新监听上
	for eventType, handlers := range s.listeners {
		for _, handler := range handlers {
			s.AddEventHandler(eventType, handler)
		}
	}

	s.client = client
	wg.Wait()
}

func (s *MessageBusService) AddEventHandler(eventType string, handler func(Event)) {
	s.listeners[eventType] = append(s.listeners[eventType], handler)

	if s.client == nil {
		// 如果client没有建立，先加到listen。

		// 上面是先把现在的handlers都加到client，才添加到client。
		// 避免这里添加一次事件，然后初始化又添加一次的双重监听
		return
	}
	s.client.On(eventType, func(event interface{}) {
		rawEvent := event.(map[string]interface{})

		var msgEvent Event
		msgEvent.Name = eventType

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

func (s *MessageBusService) PublishEvent(eventType EventType, properties map[string]string) error {
	// 等待client建立
	// 如果尝试超过3次仍然无法连接，则退出循环
	attempts := 0
	maxAttempts := 3

	for s.client == nil {
		attempts++
		if attempts > maxAttempts {
			return fmt.Errorf("无法连接到消息总线，已尝试%d次", maxAttempts)
		}
		time.Sleep(2 * time.Second)
	}

	room := "event"
	if eventType.Room != "" {
		room = eventType.Room
	}

	// 如果报错是 connection reset by peer，那么说明 properties 太大了,超过buffer了
	err := s.client.Emit(eventType.Name, room, eventType.SourceID, properties, room)
	return err
}
