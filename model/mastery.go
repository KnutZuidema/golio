package model

type Mastery struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  []string  `json:"description"`
	Image        ImageData `json:"image"`
	Ranks        int       `json:"ranks"`
	Prerequisite string    `json:"prereq"`
}
