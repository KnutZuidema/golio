package model

// StatusTranslation contains the status message content in a certain language
type StatusTranslation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}
