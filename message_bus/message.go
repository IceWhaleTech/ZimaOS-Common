package message_bus

import "time"

type EventType struct {
	Name             string         `json:"name"`
	PropertyTypeList []PropertyType `json:"propertyTypeList"`
	SourceID         string         `json:"sourceID"`
	Room             string         `json:"room"`
}

// PropertyType defines model for PropertyType.
type PropertyType struct {
	Description *string `json:"description,omitempty"`
	Example     *string `json:"example,omitempty"`
	Name        string  `json:"name"`
}

type Event struct {
	// Name event name
	Name string `json:"name"`

	// Properties event properties
	Properties map[string]string `json:"properties"`

	// SourceID associated source id
	SourceID string `json:"sourceID"`

	// Timestamp timestamp this event took place
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Uuid event uuid
	Uuid *string `json:"uuid,omitempty"`
}
