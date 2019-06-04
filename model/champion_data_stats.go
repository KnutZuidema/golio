package model

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
