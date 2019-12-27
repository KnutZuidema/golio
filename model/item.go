package model

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
