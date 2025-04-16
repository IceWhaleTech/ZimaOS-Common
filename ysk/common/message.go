package common

import (
	"github.com/IceWhaleTech/ZimaOS-Common/message_bus"
	"github.com/samber/lo"
)

const (
	SERVICENAME = "ysk"
)

// common properties
var (
	PropertyTypeMessage = message_bus.PropertyType{
		Name:        "message",
		Description: lo.ToPtr("message at different levels, typically for error"),
	}
)

// app properties
var (
	PropertyTypeCardID = message_bus.PropertyType{
		Name:        "card:id",
		Description: lo.ToPtr("card id"),
		Example:     lo.ToPtr("task:application:install"),
	}

	PropertyTypeCardBody = message_bus.PropertyType{
		Name:        "card:body",
		Description: lo.ToPtr("card body"),
		Example:     lo.ToPtr("{xxxxxx}"),
	}
)

var (
	EventTypeYSKCardUpsert = message_bus.EventType{
		SourceID: SERVICENAME,
		Name:     "ysk:card:upsert",
		Room:     "ysk",
		PropertyTypeList: []message_bus.PropertyType{
			PropertyTypeCardBody,
		},
	}

	EventTypeYSKCardDelete = message_bus.EventType{
		SourceID: SERVICENAME,
		Name:     "ysk:card:delete",
		Room:     "ysk",
		PropertyTypeList: []message_bus.PropertyType{
			PropertyTypeCardID,
		},
	}
)
