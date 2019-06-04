package model

// SpellData contains information about a champions spell
type SpellData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"abstract"`
	Tooltip     string `json:"tooltip"`
	Leveltip    struct {
		Label  []string `json:"label"`
		Effect []string `json:"effect"`
	} `json:"leveltip"`
	MaxRank      int         `json:"maxrank"`
	Cooldown     []float64   `json:"cooldown"`
	CooldownBurn string      `json:"cooldownBurn"`
	Cost         []float64   `json:"cost"`
	CostBurn     string      `json:"costBurn"`
	Effect       [][]float64 `json:"effect"`
	EffectBurn   []string    `json:"effectBurn"`
	Vars         []struct {
		Link        string  `json:"link"`
		Coefficient float64 `json:"coeff"`
		Key         string  `json:"key"`
	} `json:"vars"`
	CostType  string    `json:"costType"`
	MaxAmmo   string    `json:"maxammo"`
	Range     []float64 `json:"range"`
	RangeBurn string    `json:"rangeBurn"`
	Image     ImageData `json:"image"`
	Resource  string    `json:"resource"`
}
