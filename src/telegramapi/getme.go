package telegramapi

// GetMeT Type
type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}
