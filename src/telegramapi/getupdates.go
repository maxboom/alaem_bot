package telegramapi

// GetUpdatesT Type
type GetUpdatesT struct {
	Ok     bool                `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}
