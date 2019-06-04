package model

// StatusMessage contains information about a status message
type StatusMessage struct {
	Severity     string              `json:"severity"`
	Author       string              `json:"author"`
	CreatedAt    string              `json:"created_at"`
	Translations []StatusTranslation `json:"translations"`
	UpdatedAt    string              `json:"updated_at"`
	Content      string              `json:"content"`
	ID           string              `json:"id"`
}
