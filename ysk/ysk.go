package ysk

import (
	"context"
	"encoding/json"

	"github.com/IceWhaleTech/ZimaOS-Common/message_bus"
	"github.com/IceWhaleTech/ZimaOS-Common/ysk/common"
)

func DefineCard(ctx context.Context, cardID string) YSKCard {
	return YSKCard{}
}

func UpsertYSKCard(ctx context.Context, YSKCard YSKCard, publish func(context.Context, message_bus.EventType, map[string]string)) error {
	yskCardBodyJSON, _ := json.Marshal(YSKCard)
	publish(ctx,
		common.EventTypeYSKCardUpsert,
		map[string]string{
			common.PropertyTypeCardBody.Name: string(yskCardBodyJSON),
		},
	)
	return nil
}

func DeleteCard(ctx context.Context, cardID string, publish func(context.Context, message_bus.EventType, map[string]string)) error {
	// do something
	publish(ctx,
		common.EventTypeYSKCardDelete,
		map[string]string{
			common.PropertyTypeCardID.Name: cardID,
		})
	return nil
}
