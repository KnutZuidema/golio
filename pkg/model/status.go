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
}
