package model

// PassiveData contains information about a champions passive ability
type PassiveData struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       ImageData `json:"image"`
}
