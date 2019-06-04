package model

// Incident contains information about an incident
type Incident struct {
	Active    bool            `json:"active"`
	CreatedAt string          `json:"created_at"`
	ID        int             `json:"id"`
	Updates   []StatusMessage `json:"updates"`
}
