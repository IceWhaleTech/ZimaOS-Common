package codegen

const (
	YSKCardCardTypeLongNotice  YSKCardCardType = "long-notice"
	YSKCardCardTypeShortNotice YSKCardCardType = "short-notice"
	YSKCardCardTypeTask        YSKCardCardType = "task"
)

const (
	YSKCardRenderTypeIconTextNotice YSKCardRenderType = "icon-text-notice"
	YSKCardRenderTypeListNotice     YSKCardRenderType = "list-notice"
	YSKCardRenderTypeMarkdownNotice YSKCardRenderType = "markdown-notice"
	YSKCardRenderTypeTask           YSKCardRenderType = "task"
)

type YSKCard struct {
	CardType   YSKCardCardType   `json:"cardType"`
	Content    YSKCardContent    `json:"content"`
	Id         string            `json:"id"`
	RenderType YSKCardRenderType `json:"renderType"`
}

type YSKCardCardType string

type YSKCardRenderType string

type YSKCardContent struct {
	BodyIconWithText *YSKCardIconWithText   `json:"bodyIconWithText,omitempty"`
	BodyList         *[]YSKCardListItem     `json:"bodyList,omitempty"`
	BodyProgress     *YSKCardProgress       `json:"bodyProgress,omitempty"`
	FooterActions    *[]YSKCardFooterAction `json:"footerActions,omitempty"`
	TitleIcon        YSKCardIcon            `json:"titleIcon"`
	TitleText        string                 `json:"titleText"`
}

type YSKCardFooterAction struct {
	MessageBus YSKCardMessageBusAction `json:"messageBus"`
	Side       string                  `json:"side"`
	Style      string                  `json:"style"`
	Text       string                  `json:"text"`
}

type YSKCardIcon = string

type YSKCardIconWithText struct {
	Description string      `json:"description"`
	Icon        YSKCardIcon `json:"icon"`
}

type YSKCardList = []YSKCard

type YSKCardListItem struct {
	Description string      `json:"description"`
	Icon        YSKCardIcon `json:"icon"`
	RightText   string      `json:"rightText"`
}

type YSKCardMessageBusAction struct {
	Key     string `json:"key"`
	Payload string `json:"payload"`
}

type YSKCardProgress struct {
	Label    string `json:"label"`
	Progress int    `json:"progress"`
}
