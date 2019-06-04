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
