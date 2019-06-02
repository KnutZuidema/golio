package model

// SummonerSpell represents a summoner spell
type SummonerSpell struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Tooltip      string    `json:"tooltip"`
	MaxRank      int       `json:"maxrank"`
	Cooldown     []float64 `json:"cooldown"`
	CooldownBurn string    `json:"cooldownBurn"`
	Cost         []float64 `json:"cost"`
	CostBurn     string    `json:"costBurn"`
	Vars         []struct {
		Link        string `json:"link"`
		Coefficient int    `json:"coeff"`
		Key         string `json:"key"`
	} `json:"vars"`
	Key           string    `json:"key"`
	SummonerLevel int       `json:"summonerLevel"`
	Modes         []string  `json:"modes"`
	CostType      string    `json:"costType"`
	MaxAmmo       string    `json:"maxammo"`
	Range         []float64 `json:"range"`
	RangeBurn     string    `json:"rangeBurn"`
	Image         ImageData `json:"image"`
	Resource      string    `json:"resource"`
}
