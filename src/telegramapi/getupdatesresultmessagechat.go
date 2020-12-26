package telegramapi

// GetUpdatesResultMessageChatT Type
type GetUpdatesResultMessageChatT struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
