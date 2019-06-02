package model

// Item represents an item
type Item struct {
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
	FlatHPPoolMod                       int `json:"FlatHPPoolMod"`
	RFlatHPModPerLevel                  int `json:"rFlatHPModPerLevel"`
	FlatMPPoolMod                       int `json:"FlatMPPoolMod"`
	RFlatMPModPerLevel                  int `json:"rFlatMPModPerLevel"`
	PercentHPPoolMod                    int `json:"PercentHPPoolMod"`
	PercentMPPoolMod                    int `json:"PercentMPPoolMod"`
	FlatHPRegenMod                      int `json:"FlatHPRegenMod"`
	RFlatHPRegenModPerLevel             int `json:"rFlatHPRegenModPerLevel"`
	PercentHPRegenMod                   int `json:"PercentHPRegenMod"`
	FlatMPRegenMod                      int `json:"FlatMPRegenMod"`
	RFlatMPRegenModPerLevel             int `json:"rFlatMPRegenModPerLevel"`
	PercentMPRegenMod                   int `json:"PercentMPRegenMod"`
	FlatArmorMod                        int `json:"FlatArmorMod"`
	RFlatArmorModPerLevel               int `json:"rFlatArmorModPerLevel"`
	PercentArmorMod                     int `json:"PercentArmorMod"`
	RFlatArmorPenetrationMod            int `json:"rFlatArmorPenetrationMod"`
	RFlatArmorPenetrationModPerLevel    int `json:"rFlatArmorPenetrationModPerLevel"`
	RPercentArmorPenetrationMod         int `json:"rPercentArmorPenetrationMod"`
	RPercentArmorPenetrationModPerLevel int `json:"rPercentArmorPenetrationModPerLevel"`
	FlatPhysicalDamageMod               int `json:"FlatPhysicalDamageMod"`
	RFlatPhysicalDamageModPerLevel      int `json:"rFlatPhysicalDamageModPerLevel"`
	PercentPhysicalDamageMod            int `json:"PercentPhysicalDamageMod"`
	FlatMagicDamageMod                  int `json:"FlatMagicDamageMod"`
	RFlatMagicDamageModPerLevel         int `json:"rFlatMagicDamageModPerLevel"`
	PercentMagicDamageMod               int `json:"PercentMagicDamageMod"`
	FlatMovementSpeedMod                int `json:"FlatMovementSpeedMod"`
	RFlatMovementSpeedModPerLevel       int `json:"rFlatMovementSpeedModPerLevel"`
	PercentMovementSpeedMod             int `json:"PercentMovementSpeedMod"`
	RPercentMovementSpeedModPerLevel    int `json:"rPercentMovementSpeedModPerLevel"`
	FlatAttackSpeedMod                  int `json:"FlatAttackSpeedMod"`
	PercentAttackSpeedMod               int `json:"PercentAttackSpeedMod"`
	RPercentAttackSpeedModPerLevel      int `json:"rPercentAttackSpeedModPerLevel"`
	RFlatDodgeMod                       int `json:"rFlatDodgeMod"`
	RFlatDodgeModPerLevel               int `json:"rFlatDodgeModPerLevel"`
	PercentDodgeMod                     int `json:"PercentDodgeMod"`
	FlatCritChanceMod                   int `json:"FlatCritChanceMod"`
	RFlatCritChanceModPerLevel          int `json:"rFlatCritChanceModPerLevel"`
	PercentCritChanceMod                int `json:"PercentCritChanceMod"`
	FlatCritDamageMod                   int `json:"FlatCritDamageMod"`
	RFlatCritDamageModPerLevel          int `json:"rFlatCritDamageModPerLevel"`
	PercentCritDamageMod                int `json:"PercentCritDamageMod"`
	FlatBlockMod                        int `json:"FlatBlockMod"`
	PercentBlockMod                     int `json:"PercentBlockMod"`
	FlatSpellBlockMod                   int `json:"FlatSpellBlockMod"`
	RFlatSpellBlockModPerLevel          int `json:"rFlatSpellBlockModPerLevel"`
	PercentSpellBlockMod                int `json:"PercentSpellBlockMod"`
	FlatEXPBonus                        int `json:"FlatEXPBonus"`
	PercentEXPBonus                     int `json:"PercentEXPBonus"`
	RPercentCooldownMod                 int `json:"rPercentCooldownMod"`
	RPercentCooldownModPerLevel         int `json:"rPercentCooldownModPerLevel"`
	RFlatTimeDeadMod                    int `json:"rFlatTimeDeadMod"`
	RFlatTimeDeadModPerLevel            int `json:"rFlatTimeDeadModPerLevel"`
	RPercentTimeDeadMod                 int `json:"rPercentTimeDeadMod"`
	RPercentTimeDeadModPerLevel         int `json:"rPercentTimeDeadModPerLevel"`
	RFlatGoldPer10Mod                   int `json:"rFlatGoldPer10Mod"`
	RFlatMagicPenetrationMod            int `json:"rFlatMagicPenetrationMod"`
	RFlatMagicPenetrationModPerLevel    int `json:"rFlatMagicPenetrationModPerLevel"`
	RPercentMagicPenetrationMod         int `json:"rPercentMagicPenetrationMod"`
	RPercentMagicPenetrationModPerLevel int `json:"rPercentMagicPenetrationModPerLevel"`
	FlatEnergyRegenMod                  int `json:"FlatEnergyRegenMod"`
	RFlatEnergyRegenModPerLevel         int `json:"rFlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                   int `json:"FlatEnergyPoolMod"`
	RFlatEnergyModPerLevel              int `json:"rFlatEnergyModPerLevel"`
	PercentLifeStealMod                 int `json:"PercentLifeStealMod"`
	PercentSpellVampMod                 int `json:"PercentSpellVampMod"`
}
