package model

// RecommendedItemSet is a set of items used in a recommended build
type RecommendedItemSet struct {
	Type                string            `json:"type"`
	RecMath             bool              `json:"recMath"`
	RecSteps            bool              `json:"recSteps"`
	MinSummonerLevel    int               `json:"minSummonerLevel"`
	MaxSummonerLevel    int               `json:"maxSummonerLevel"`
	ShowIfSummonerSpell string            `json:"showIfSummonerSpell"`
	HideIfSummonerSpell string            `json:"hideIfSummonerSpell"`
	Items               []RecommendedItem `json:"items"`
}
