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

// ChampionDataInfo contains information about the playstyle of a champion
type ChampionDataInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

// ImageData contains information about an image
type ImageData struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}

// ChampionDataStats contains information about the stats of a champion
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

// ChampionDataExtended contains additional data about a champion
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

// SkinData contains information about a skin for a champion
type SkinData struct {
	ID      string `json:"id"`
	Num     int    `json:"num"`
	Name    string `json:"name"`
	Chromas bool   `json:"chromas"`
}

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

// PassiveData contains information about a champions passive ability
type PassiveData struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       ImageData `json:"image"`
}

// RecommendedItemData a build recommended for a champion
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

// RecommendedItem represents an item in a recommended set
type RecommendedItem struct {
	ID        string `json:"id"`
	Count     int    `json:"count"`
	HideCount bool   `json:"hideCount"`
}

// Item represents an item
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Rune struct {
		IsRune bool   `json:"isrune"`
		Tier   int    `json:"tier"`
		Type   string `json:"type"`
	} `json:"rune"`
	Gold struct {
		Base        int  `json:"base"`
		Total       int  `json:"total"`
		Sell        int  `json:"sell"`
		Purchasable bool `json:"purchasable"`
	} `json:"gold"`
	Group            string          `json:"group"`
	Description      string          `json:"description"`
	Colloqial        string          `json:"colloq"`
	Plaintext        string          `json:"plaintext"`
	Consumed         bool            `json:"consumed"`
	Stacks           int             `json:"stacks"`
	Depth            int             `json:"depth"`
	ConsumeOnFull    bool            `json:"consumeOnFull"`
	From             []string        `json:"from"`
	Into             []string        `json:"into"`
	SpecialRecipe    int             `json:"specialRecipe"`
	InStore          bool            `json:"inStore"`
	HideFromAll      bool            `json:"hideFromAll"`
	RequiredChampion string          `json:"requiredChampion"`
	Stats            ItemStats       `json:"stats"`
	Tags             []string        `json:"tags"`
	Maps             map[string]bool `json:"maps"`
}

// ItemStats contains information about the stats of an item
type ItemStats struct {
	FlatHPPoolMod                       float64 `json:"FlatHPPoolMod"`
	RFlatHPModPerLevel                  float64 `json:"rFlatHPModPerLevel"`
	FlatMPPoolMod                       float64 `json:"FlatMPPoolMod"`
	RFlatMPModPerLevel                  float64 `json:"rFlatMPModPerLevel"`
	PercentHPPoolMod                    float64 `json:"PercentHPPoolMod"`
	PercentMPPoolMod                    float64 `json:"PercentMPPoolMod"`
	FlatHPRegenMod                      float64 `json:"FlatHPRegenMod"`
	RFlatHPRegenModPerLevel             float64 `json:"rFlatHPRegenModPerLevel"`
	PercentHPRegenMod                   float64 `json:"PercentHPRegenMod"`
	FlatMPRegenMod                      float64 `json:"FlatMPRegenMod"`
	RFlatMPRegenModPerLevel             float64 `json:"rFlatMPRegenModPerLevel"`
	PercentMPRegenMod                   float64 `json:"PercentMPRegenMod"`
	FlatArmorMod                        float64 `json:"FlatArmorMod"`
	RFlatArmorModPerLevel               float64 `json:"rFlatArmorModPerLevel"`
	PercentArmorMod                     float64 `json:"PercentArmorMod"`
	RFlatArmorPenetrationMod            float64 `json:"rFlatArmorPenetrationMod"`
	RFlatArmorPenetrationModPerLevel    float64 `json:"rFlatArmorPenetrationModPerLevel"`
	RPercentArmorPenetrationMod         float64 `json:"rPercentArmorPenetrationMod"`
	RPercentArmorPenetrationModPerLevel float64 `json:"rPercentArmorPenetrationModPerLevel"`
	FlatPhysicalDamageMod               float64 `json:"FlatPhysicalDamageMod"`
	RFlatPhysicalDamageModPerLevel      float64 `json:"rFlatPhysicalDamageModPerLevel"`
	PercentPhysicalDamageMod            float64 `json:"PercentPhysicalDamageMod"`
	FlatMagicDamageMod                  float64 `json:"FlatMagicDamageMod"`
	RFlatMagicDamageModPerLevel         float64 `json:"rFlatMagicDamageModPerLevel"`
	PercentMagicDamageMod               float64 `json:"PercentMagicDamageMod"`
	FlatMovementSpeedMod                float64 `json:"FlatMovementSpeedMod"`
	RFlatMovementSpeedModPerLevel       float64 `json:"rFlatMovementSpeedModPerLevel"`
	PercentMovementSpeedMod             float64 `json:"PercentMovementSpeedMod"`
	RPercentMovementSpeedModPerLevel    float64 `json:"rPercentMovementSpeedModPerLevel"`
	FlatAttackSpeedMod                  float64 `json:"FlatAttackSpeedMod"`
	PercentAttackSpeedMod               float64 `json:"PercentAttackSpeedMod"`
	RPercentAttackSpeedModPerLevel      float64 `json:"rPercentAttackSpeedModPerLevel"`
	RFlatDodgeMod                       float64 `json:"rFlatDodgeMod"`
	RFlatDodgeModPerLevel               float64 `json:"rFlatDodgeModPerLevel"`
	PercentDodgeMod                     float64 `json:"PercentDodgeMod"`
	FlatCritChanceMod                   float64 `json:"FlatCritChanceMod"`
	RFlatCritChanceModPerLevel          float64 `json:"rFlatCritChanceModPerLevel"`
	PercentCritChanceMod                float64 `json:"PercentCritChanceMod"`
	FlatCritDamageMod                   float64 `json:"FlatCritDamageMod"`
	RFlatCritDamageModPerLevel          float64 `json:"rFlatCritDamageModPerLevel"`
	PercentCritDamageMod                float64 `json:"PercentCritDamageMod"`
	FlatBlockMod                        float64 `json:"FlatBlockMod"`
	PercentBlockMod                     float64 `json:"PercentBlockMod"`
	FlatSpellBlockMod                   float64 `json:"FlatSpellBlockMod"`
	RFlatSpellBlockModPerLevel          float64 `json:"rFlatSpellBlockModPerLevel"`
	PercentSpellBlockMod                float64 `json:"PercentSpellBlockMod"`
	FlatEXPBonus                        float64 `json:"FlatEXPBonus"`
	PercentEXPBonus                     float64 `json:"PercentEXPBonus"`
	RPercentCooldownMod                 float64 `json:"rPercentCooldownMod"`
	RPercentCooldownModPerLevel         float64 `json:"rPercentCooldownModPerLevel"`
	RFlatTimeDeadMod                    float64 `json:"rFlatTimeDeadMod"`
	RFlatTimeDeadModPerLevel            float64 `json:"rFlatTimeDeadModPerLevel"`
	RPercentTimeDeadMod                 float64 `json:"rPercentTimeDeadMod"`
	RPercentTimeDeadModPerLevel         float64 `json:"rPercentTimeDeadModPerLevel"`
	RFlatGoldPer10Mod                   float64 `json:"rFlatGoldPer10Mod"`
	RFlatMagicPenetrationMod            float64 `json:"rFlatMagicPenetrationMod"`
	RFlatMagicPenetrationModPerLevel    float64 `json:"rFlatMagicPenetrationModPerLevel"`
	RPercentMagicPenetrationMod         float64 `json:"rPercentMagicPenetrationMod"`
	RPercentMagicPenetrationModPerLevel float64 `json:"rPercentMagicPenetrationModPerLevel"`
	FlatEnergyRegenMod                  float64 `json:"FlatEnergyRegenMod"`
	RFlatEnergyRegenModPerLevel         float64 `json:"rFlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                   float64 `json:"FlatEnergyPoolMod"`
	RFlatEnergyModPerLevel              float64 `json:"rFlatEnergyModPerLevel"`
	PercentLifeStealMod                 float64 `json:"PercentLifeStealMod"`
	PercentSpellVampMod                 float64 `json:"PercentSpellVampMod"`
}

// Mastery represents an old mastery
type Mastery struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  []string  `json:"description"`
	Image        ImageData `json:"image"`
	Ranks        int       `json:"ranks"`
	Prerequisite string    `json:"prereq"`
}

// ProfileIcon represents a profile icon
type ProfileIcon struct {
	ID    int       `json:"id"`
	Image ImageData `json:"image"`
}

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
		Link        string      `json:"link"`
		Coefficient interface{} `json:"coeff"`
		Key         string      `json:"key"`
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
