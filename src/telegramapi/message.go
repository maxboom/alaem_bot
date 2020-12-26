package telegramapi

// MessageT Type
type MessageT struct {
	MessageID int                          `json:"message_id"`
	From      GetUpdatesResultMessageFromT `json:"from"`
	Chat      GetUpdatesResultMessageChatT `json:"chat"`
	Date      int                          `json:"date"`
	Text      string                       `json:"text"`
}
