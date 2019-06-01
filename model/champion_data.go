package model

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

type ChampionDataExtended struct {
	ChampionData
	Skins            []SkinData            `json:"skins"`
	Lore             string                `json:"lore"`
	AllyTips         []string              `json:"allytips"`
	EnemyTips        []string              `json:"enemytips"`
	Spells           []SpellData           `json:"spells"`
	Passive          PassiveData           `json:"passive"`
	RecommendedItems []RecommendedItemData `json:"recommended"`
}

type SkinData struct {
	ID      string `json:"id"`
	Num     int    `json:"num"`
	Name    string `json:"name"`
	Chromas bool   `json:"chromas"`
}

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

type PassiveData struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       ImageData `json:"image"`
}

type RecommendedItemData struct {
	Champion      string               `json:"champion"`
	Title         string               `json:"title"`
	Map           string               `json:"map"`
	Mode          string               `json:"mode"`
	CustomTag     string               `json:"customTag"`
	SortRank      int                  `json:"sortrank"`
	ExtensionPage bool                 `json:"extensionPage"`
	CustomPanel   interface{}          `json:"customPanel"`
	Blocks        []RecommendedItemSet `json:"blocks"`
}

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

type RecommendedItem struct {
	ID        string `json:"id"`
	Count     int    `json:"count"`
	HideCount bool   `json:"hideCount"`
}

type ChampionDataInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

type ImageData struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

type ChampionDataStats struct {
	HealthPoints                    float64 `json:"hp"`
	HealthPointsPerLevel            float64 `json:"hpperlevel"`
	ManaPoints                      float64 `json:"mp"`
	ManaPointsPerLevel              float64 `json:"mpperlevel"`
	MovementSpeed                   float64 `json:"movespeed"`
	Armor                           float64 `json:"armor"`
	ArmorPerLevel                   float64 `json:"armorperlevel"`
	SpellBlock                      float64 `json:"spellblock"`
	SpellBlockPerLevel              float64 `json:"spellblockperlevel"`
	AttackRange                     float64 `json:"attackrange"`
	HealthPointRegeneration         float64 `json:"hpregen"`
	HealthPointRegenerationPerLevel float64 `json:"hpregenperlevel"`
	ManaPointRegeneration           float64 `json:"mpregen"`
	ManaPointRegenerationPerLevel   float64 `json:"mpregenperlevel"`
	CriticalStrikeChance            float64 `json:"crit"`
	CriticalStrikeChancePerLevel    float64 `json:"critperlevel"`
	AttackDamage                    float64 `json:"attackdamage"`
	AttackDamagePerLevel            float64 `json:"attackdamageperlevel"`
	AttackSpeedOffset               float64 `json:"attackspeedoffset"`
	AttackSpeedPerLevel             float64 `json:"attackspeedperlevel"`
}
