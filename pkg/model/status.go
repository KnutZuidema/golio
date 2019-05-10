package model

type Status struct {
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	Hostname  string    `json:"hostname"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
	Locales   []string  `json:"locales"`
}

type Service struct {
	Status    string     `json:"status"`
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
}

type Incident struct {
	Active    bool            `json:"active"`
	CreatedAt string          `json:"created_at"`
	ID        int             `json:"id"`
	Updates   []StatusMessage `json:"updates"`
}

type StatusMessage struct {
	Severity     string              `json:"severity"`
	Author       string              `json:"author"`
	CreatedAt    string              `json:"created_at"`
	Translations []StatusTranslation `json:"translations"`
	UpdatedAt    string              `json:"updated_at"`
	Content      string              `json:"content"`
	ID           string              `json:"id"`
}

type StatusTranslation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}
