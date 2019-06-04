package model

// ParticipantStats contains stats of a participant in a game
type ParticipantStats struct {
	FirstBloodAssist            bool `json:"firstBloodAssist"`
	VisionScore                 int  `json:"visionScore"`
	MagicDamageDealtToChampions int  `json:"magicDamageDealtToChampions"`
	DamageDealtToObjectives     int  `json:"damageDealtToObjectives"`
	TotalTimeCrowdControlDealt  int  `json:"totalTimeCrowdControlDealt"`
	IntestTimeSpentLiving       int  `json:"intestTimeSpentLiving"`
	// Post game rune stats.
	Perk1Var1 int `json:"perk1Var1"`
	// Post game rune stats.
	Perk1Var3 int `json:"perk1Var3"`
	// Post game rune stats.
	Perk1Var2   int `json:"perk1Var2"`
	TripleKills int `json:"tripleKills"`
	// Post game rune stats.
	Perk3Var3            int `json:"perk3Var3"`
	NodeNeutralizeAssist int `json:"nodeNeutralizeAssist"`
	// Post game rune stats.
	Perk3Var2    int `json:"perk3Var2"`
	PlayerScore9 int `json:"playerScore9"`
	PlayerScore8 int `json:"playerScore8"`
	Kills        int `json:"kills"`
	PlayerScore1 int `json:"playerScore1"`
	PlayerScore0 int `json:"playerScore0"`
	PlayerScore3 int `json:"playerScore3"`
	PlayerScore2 int `json:"playerScore2"`
	PlayerScore5 int `json:"playerScore5"`
	PlayerScore4 int `json:"playerScore4"`
	PlayerScore7 int `json:"playerScore7"`
	PlayerScore6 int `json:"playerScore6"`
	// Post game rune stats.
	Perk5Var1 int `json:"perk5Var1"`
	// Post game rune stats.
	Perk5Var3 int `json:"perk5Var3"`
	// Post game rune stats.
	Perk5Var2                      int `json:"perk5Var2"`
	TotalScoreRank                 int `json:"totalScoreRank"`
	NeutralMinionsKilled           int `json:"neutralMinionsKilled"`
	DamageDealtToTurrets           int `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions int `json:"physicalDamageDealtToChampions"`
	NodeCapture                    int `json:"nodeCapture"`
	LargestMultiKill               int `json:"largestMultiKill"`
	// Post game rune stats.
	Perk2Var2 int `json:"perk2Var2"`
	// Post game rune stats.
	Perk2Var3        int `json:"perk2Var3"`
	TotalUnitsHealed int `json:"totalUnitsHealed"`
	// Post game rune stats.
	Perk2Var1 int `json:"perk2Var1"`
	// Post game rune stats.
	Perk4Var1 int `json:"perk4Var1"`
	// Post game rune stats.
	Perk4Var2 int `json:"perk4Var2"`
	// Post game rune stats.
	Perk4Var3                      int `json:"perk4Var3"`
	WardsKilled                    int `json:"wardsKilled"`
	LargestCriticalStrike          int `json:"largestCriticalStrike"`
	LargestKillingSpree            int `json:"largestKillingSpree"`
	QuadraKills                    int `json:"quadraKills"`
	TeamObjective                  int `json:"teamObjective"`
	MagicDamageDealt               int `json:"magicDamageDealt"`
	Item2                          int `json:"item2"`
	Item3                          int `json:"item3"`
	Item0                          int `json:"item0"`
	NeutralMinionsKilledTeamJungle int `json:"neutralMinionsKilledTeamJungle"`
	Item6                          int `json:"item6"`
	Item4                          int `json:"item4"`
	Item5                          int `json:"item5"`
	// Primary path rune.
	Perk1 int `json:"perk1"`
	// Primary path keystone rune.
	Perk0 int `json:"perk0"`
	// Primary path rune.
	Perk3 int `json:"perk3"`
	// Primary path rune.
	Perk2 int `json:"perk2"`
	// Secondary path rune.
	Perk5 int `json:"perk5"`
	// Secondary path rune.
	Perk4 int `json:"perk4"`
	// Post game rune stats.
	Perk3Var1           int  `json:"perk3Var1"`
	DamageSelfMitigated int  `json:"damageSelfMitigated"`
	MagicalDamageTaken  int  `json:"magicalDamageTaken"`
	FirstInhibitorKill  bool `json:"firstInhibitorKill"`
	TrueDamageTaken     int  `json:"trueDamageTaken"`
	NodeNeutralize      int  `json:"nodeNeutralize"`
	Assists             int  `json:"assists"`
	CombatPlayerScore   int  `json:"combatPlayerScore"`
	// Primary rune path
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	GoldSpent                       int  `json:"goldSpent"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	ParticipantID                   int  `json:"participantId"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	Win                             bool `json:"win"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	Item1                           int  `json:"item1"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	Deaths                          int  `json:"deaths"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	// Secondary rune path
	PerkSubStyle               int  `json:"perkSubStyle"`
	TurretKills                int  `json:"turretKills"`
	FirstBloodKill             bool `json:"firstBloodKill"`
	TrueDamageDealtToChampions int  `json:"trueDamageDealtToChampions"`
	GoldEarned                 int  `json:"goldEarned"`
	KillingSprees              int  `json:"killingSprees"`
	UnrealKills                int  `json:"unrealKills"`
	AltarsCaptured             int  `json:"altarsCaptured"`
	FirstTowerAssist           bool `json:"firstTowerAssist"`
	FirstTowerKill             bool `json:"firstTowerKill"`
	ChampLevel                 int  `json:"champLevel"`
	DoubleKills                int  `json:"doubleKills"`
	NodeCaptureAssist          int  `json:"nodeCaptureAssist"`
	InhibitorKills             int  `json:"inhibitorKills"`
	FirstInhibitorAssist       bool `json:"firstInhibitorAssist"`
	// Post game rune stats.
	Perk0Var1 int `json:"perk0Var1"`
	// Post game rune stats.
	Perk0Var2 int `json:"perk0Var2"`
	// Post game rune stats.
	Perk0Var3               int `json:"perk0Var3"`
	VisionWardsBoughtInGame int `json:"visionWardsBoughtInGame"`
	AltarsNeutralized       int `json:"altarsNeutralized"`
	PentaKills              int `json:"pentaKills"`
	TotalHeal               int `json:"totalHeal"`
	TotalMinionsKilled      int `json:"totalMinionsKilled"`
	TimeCCingOthers         int `json:"timeCCingOthers"`
}
