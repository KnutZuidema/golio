package model

// ChampionData contains information about a champion
type ChampionData struct {
	Version string            `json:"version"`
	ID      string            `json:"id"`
	Key     string            `json:"key"`
	Name    string            `json:"name"`
	Title   string            `json:"title"`
	Blurb   string            `json:"blurb"`
	Info    ChampionDataInfo  `json:"info"`
	Image   ImageData         `json:"image"`
	Tags    []string          `json:"tags"`
	Partype string            `json:"partype"`
	Stats   ChampionDataStats `json:"stats"`
}
