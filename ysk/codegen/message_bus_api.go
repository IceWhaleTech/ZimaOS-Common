package codegen

import (
	"time"
)

const (
	Access_tokenScopes = "access_token.Scopes"
)

// Defines values for YSKCardCardType.
const (
	YSKCardCardTypeLongNotice  YSKCardCardType = "long-notice"
	YSKCardCardTypeShortNotice YSKCardCardType = "short-notice"
	YSKCardCardTypeTask        YSKCardCardType = "task"
)

// Defines values for YSKCardRenderType.
const (
	YSKCardRenderTypeIconTextNotice YSKCardRenderType = "icon-text-notice"
	YSKCardRenderTypeListNotice     YSKCardRenderType = "list-notice"
	YSKCardRenderTypeMarkdownNotice YSKCardRenderType = "markdown-notice"
	YSKCardRenderTypeTask           YSKCardRenderType = "task"
)

// Action defines model for Action.
type Action struct {
	// Name action name
	Name string `json:"name"`

	// Properties event properties
	Properties map[string]string `json:"properties"`

	// SourceID associated source id
	SourceID string `json:"sourceID"`

	// Timestamp timestamp this action took place
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// ActionType defines model for ActionType.
type ActionType struct {
	// Name action name
	//
	// (there is no naming convention for action names, but it is recommended to name each as structural and descriptive as possible)
	Name             string         `json:"name"`
	PropertyTypeList []PropertyType `json:"propertyTypeList"`

	// SourceID action source id to identify where the action will take
	SourceID string `json:"sourceID"`
}

// BaseResponse defines model for BaseResponse.
type BaseResponse struct {
	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// Event defines model for Event.
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

// EventType defines model for EventType.
type EventType struct {
	// Name event name
	//
	// (there is no naming convention for event names, but it is recommended to name each as structural and descriptive as possible)
	Name             string         `json:"name"`
	PropertyTypeList []PropertyType `json:"propertyTypeList"`

	// SourceID event source id to identify where the event comes from
	SourceID string `json:"sourceID"`
}

// PropertyType defines model for PropertyType.
type PropertyType struct {
	Description *string `json:"description,omitempty"`
	Example     *string `json:"example,omitempty"`

	// Name property name
	//
	// > It is recommended for a property name to be as descriptive as possible. One option is to prefix with a namespace.
	// > - If the property is source specific, prefix with source ID. For example, `local-storage:vendor`
	// > - Otherwise, prefix with `common:`. For example, `common:email`
	// >
	// > Some bad examples are `id`, `avail`, `blk`...which can be ambiguous and confusing.
	Name string `json:"name"`
}

// YSKCard defines model for YSKCard.
type YSKCard struct {
	CardType   YSKCardCardType   `json:"cardType"`
	Content    YSKCardContent    `json:"content"`
	Id         string            `json:"id"`
	RenderType YSKCardRenderType `json:"renderType"`
}

// YSKCardCardType defines model for YSKCard.CardType.
type YSKCardCardType string

// YSKCardRenderType defines model for YSKCard.RenderType.
type YSKCardRenderType string

// YSKCardContent defines model for YSKCardContent.
type YSKCardContent struct {
	BodyIconWithText *YSKCardIconWithText   `json:"bodyIconWithText,omitempty"`
	BodyList         *[]YSKCardListItem     `json:"bodyList,omitempty"`
	BodyProgress     *YSKCardProgress       `json:"bodyProgress,omitempty"`
	FooterActions    *[]YSKCardFooterAction `json:"footerActions,omitempty"`
	TitleIcon        YSKCardIcon            `json:"titleIcon"`
	TitleText        string                 `json:"titleText"`
}

// YSKCardFooterAction defines model for YSKCardFooterAction.
type YSKCardFooterAction struct {
	MessageBus YSKCardMessageBusAction `json:"messageBus"`
	Side       string                  `json:"side"`
	Style      string                  `json:"style"`
	Text       string                  `json:"text"`
}

// YSKCardIcon defines model for YSKCardIcon.
type YSKCardIcon = string

// YSKCardIconWithText defines model for YSKCardIconWithText.
type YSKCardIconWithText struct {
	Description string      `json:"description"`
	Icon        YSKCardIcon `json:"icon"`
}

// YSKCardList defines model for YSKCardList.
type YSKCardList = []YSKCard

// YSKCardListItem defines model for YSKCardListItem.
type YSKCardListItem struct {
	Description string      `json:"description"`
	Icon        YSKCardIcon `json:"icon"`
	RightText   string      `json:"rightText"`
}

// YSKCardMessageBusAction defines model for YSKCardMessageBusAction.
type YSKCardMessageBusAction struct {
	Key     string `json:"key"`
	Payload string `json:"payload"`
}

// YSKCardProgress defines model for YSKCardProgress.
type YSKCardProgress struct {
	Label    string `json:"label"`
	Progress int    `json:"progress"`
}

// ActionName defines model for ActionName.
type ActionName = string

// ActionNames defines model for ActionNames.
type ActionNames = []string

// EventName defines model for EventName.
type EventName = string

// EventNames defines model for EventNames.
type EventNames = []string

// SourceID defines model for SourceID.
type SourceID = string

// GetActionTypeOK defines model for GetActionTypeOK.
type GetActionTypeOK = ActionType

// GetActionTypesOK defines model for GetActionTypesOK.
type GetActionTypesOK = []ActionType

// GetEventTypeOK defines model for GetEventTypeOK.
type GetEventTypeOK = EventType

// GetEventTypesOK defines model for GetEventTypesOK.
type GetEventTypesOK = []EventType

// PublishEventOK defines model for PublishEventOK.
type PublishEventOK = Event

// ResponseBadRequest defines model for ResponseBadRequest.
type ResponseBadRequest = BaseResponse

// ResponseConflict defines model for ResponseConflict.
type ResponseConflict = BaseResponse

// ResponseGetYSKCardListOK defines model for ResponseGetYSKCardListOK.
type ResponseGetYSKCardListOK struct {
	Data *YSKCardList `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// ResponseInternalServerError defines model for ResponseInternalServerError.
type ResponseInternalServerError = BaseResponse

// ResponseNotFound defines model for ResponseNotFound.
type ResponseNotFound = BaseResponse

// ResponseOK defines model for ResponseOK.
type ResponseOK = BaseResponse

// TriggerActionOK defines model for TriggerActionOK.
type TriggerActionOK = Action

// PublishEvent event properties
type PublishEvent map[string]string

// RegisterActionTypes defines model for RegisterActionTypes.
type RegisterActionTypes = []ActionType

// RegisterEventTypes defines model for RegisterEventTypes.
type RegisterEventTypes = []EventType

// TriggerAction action properties
type TriggerAction map[string]string

// SubscribeActionWSParams defines parameters for SubscribeActionWS.
type SubscribeActionWSParams struct {
	Names *ActionNames `form:"names,omitempty" json:"names,omitempty"`
}

// TriggerActionJSONBody defines parameters for TriggerAction.
type TriggerActionJSONBody map[string]string

// RegisterActionTypesJSONBody defines parameters for RegisterActionTypes.
type RegisterActionTypesJSONBody = []ActionType

// SubscribeEventWSParams defines parameters for SubscribeEventWS.
type SubscribeEventWSParams struct {
	Names *EventNames `form:"names,omitempty" json:"names,omitempty"`
}

// PublishEventJSONBody defines parameters for PublishEvent.
type PublishEventJSONBody map[string]string

// RegisterEventTypesJSONBody defines parameters for RegisterEventTypes.
type RegisterEventTypesJSONBody = []EventType

// TriggerActionJSONRequestBody defines body for TriggerAction for application/json ContentType.
type TriggerActionJSONRequestBody TriggerActionJSONBody

// RegisterActionTypesJSONRequestBody defines body for RegisterActionTypes for application/json ContentType.
type RegisterActionTypesJSONRequestBody = RegisterActionTypesJSONBody

// PublishEventJSONRequestBody defines body for PublishEvent for application/json ContentType.
type PublishEventJSONRequestBody PublishEventJSONBody

// RegisterEventTypesJSONRequestBody defines body for RegisterEventTypes for application/json ContentType.
type RegisterEventTypesJSONRequestBody = RegisterEventTypesJSONBody

// ServerInterface represents all server handlers.
