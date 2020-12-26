package telegramapi

// GetUpdatesResultT Type
type GetUpdatesResultT struct {
	UpdateID int                `json:"update_id"`
	Message  GetUpdatesMessageT `json:"message,omitempty"`
}
