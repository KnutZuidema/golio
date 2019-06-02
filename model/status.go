package model

// Status contains information about all services in a certain region
type Status struct {
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	Hostname  string    `json:"hostname"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
	Locales   []string  `json:"locales"`
}

// Service is a service provided by Riot with its status
type Service struct {
	Status    string     `json:"status"`
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
}

// Incident contains information about an incident
type Incident struct {
	Active    bool            `json:"active"`
	CreatedAt string          `json:"created_at"`
	ID        int             `json:"id"`
	Updates   []StatusMessage `json:"updates"`
}

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

// StatusTranslation contains the status message content in a certain language
type StatusTranslation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}
