package model

// MiniSeries represents a mini series when playing to ascend to the next ranked tier
type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}
