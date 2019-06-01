package model

type Match struct {
	// Please refer to the Game Constants documentation.
	SeasonID int `json:"seasonId"`
	// Please refer to the Game Constants documentation.
	QueueID int `json:"queueId"`
	GameID  int `json:"gameId"`
	// Participant identity information.
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	// The major.minor version typically indicates the patch the match was played on.
	GameVersion string `json:"gameVersion"`
	// Platform where the match was played.
	PlatformID string `json:"platformId"`
	// Please refer to the Game Constants documentation.
	GameMode string `json:"gameMode"`
	// Please refer to the Game Constants documentation.
	MapID int `json:"mapId"`
	// Please refer to the Game Constants documentation.
	GameType string `json:"gameType"`
	// Team information.
	Teams []TeamStats `json:"teams"`
	// Participant information.
	Participants []Participant `json:"participants"`
	// Match duration in seconds.
	GameDuration int `json:"gameDuration"`
	// Designates the timestamp when champion select ended and the loading screen appeared, NOT when the game timer was at 0:00.
	GameCreation int `json:"gameCreation"`
}

type ParticipantIdentity struct {
	// Player information.
	Player        Player `json:"player"`
	ParticipantID int    `json:"participantId"`
}

type Player struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryUri   string `json:"matchHistoryUri"`
	// Original platformId.
	PlatformID string `json:"platformId"`
	// Player's current accountID (Encrypted)
	CurrentAccountID string `json:"currentAccountId"`
	ProfileIcon      int    `json:"profileIcon"`
	// Player's summonerID (Encrypted)
	SummonerID string `json:"summonerId"`
	// Player's original accountID (Encrypted)
	AccountID string `json:"accountId"`
}

type TeamStats struct {
	// Flag indicating whether or not the team scored the first Dragon kill.
	FirstDragon bool `json:"firstDragon"`
	// Flag indicating whether or not the team destroyed the first inhibitor.
	FirstInhibitor bool `json:"firstInhibitor"`
	// If match queueID has a draft, contains banned champion data, otherwise empty.
	Bans []TeamBans `json:"bans"`
	// Number of times the team killed Baron.
	BaronKills int `json:"baronKills"`
	// Flag indicating whether or not the team scored the first Rift Herald kill.
	FirstRiftHerald bool `json:"firstRiftHerald"`
	// Flag indicating whether or not the team scored the first Baron kill.
	FirstBaron bool `json:"firstBaron"`
	// Number of times the team killed Rift Herald.
	RiftHeraldKills int `json:"riftHeraldKills"`
	// Flag indicating whether or not the team scored the first blood.
	FirstBlood bool `json:"firstBlood"`
	// 100 for blue side. 200 for red side.
	TeamID int `json:"teamId"`
	// Flag indicating whether or not the team destroyed the first tower.
	FirstTower bool `json:"firstTower"`
	// Number of times the team killed Vilemaw.
	VilemawKills int `json:"vilemawKills"`
	// Number of inhibitors the team destroyed.
	InhibitorKills int `json:"inhibitorKills"`
	// Number of towers the team destroyed.
	TowerKills int `json:"towerKills"`
	// For Dominion matches, specifies the points the team had at game end.
	DominionVictoryScore int `json:"dominionVictoryScore"`
	// string indicating whether or not the team won. There are only two values visibile in public match history. (Legal values: Fail, Win)
	Win string `json:"win"`
	// Number of times the team killed Dragon.
	DragonKills int `json:"dragonKills"`
}

type TeamBans struct {
	// Turn during which the champion was banned.
	PickTurn int `json:"pickTurn"`
	// Banned championId.
	ChampionID int `json:"championId"`
}

type Participant struct {
	// Participant statistics.
	Stats         ParticipantStats `json:"stats"`
	ParticipantID int              `json:"participantId"`
	// List of legacy Rune information. Not included for matches played with Runes Reforged.
	Runes []Rune `json:"runes"`
	// Participant timeline data.
	Timeline ParticipantTimeline `json:"timeline"`
	// 100 for blue side. 200 for red side.
	TeamID int `json:"teamId"`
	// Second Summoner Spell id.
	Spell2ID int `json:"spell2Id"`
	// List of legacy Mastery information. Not included for matches played with Runes Reforged.
	Masteries []LegacyMastery `json:"masteries"`
	// Highest ranked tier achieved for the previous season in a specific subset of queueIds, if any, otherwise null. Used to display border in game loading screen. Please refer to the Ranked Info documentation. (Legal values: CHALLENGER, MASTER, DIAMOND, PLATINUM, GOLD, SILVER, BRONZE, UNRANKED)
	HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
	// First Summoner Spell id.
	Spell1ID   int `json:"spell1Id"`
	ChampionID int `json:"championId"`
}

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

type Rune struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

type ParticipantTimeline struct {
	// Participant's calculated lane. MID and BOT are legacy values. (Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM)
	Lane          string `json:"lane"`
	ParticipantID int    `json:"participantId"`
	// Creep score difference versus the calculated lane opponent(s) for a specified period.
	CsDiffPerMinDeltas map[string]float64 `json:"csDiffPerMinDeltas"`
	// Gold for a specified period.
	GoldPerMinDeltas map[string]float64 `json:"goldPerMinDeltas"`
	// Experience difference versus the calculated lane opponent(s) for a specified period.
	XpDiffPerMinDeltas map[string]float64 `json:"xpDiffPerMinDeltas"`
	// Creeps for a specified period.
	CreepsPerMinDeltas map[string]float64 `json:"creepsPerMinDeltas"`
	// Experience change for a specified period.
	XpPerMinDeltas map[string]float64 `json:"xpPerMinDeltas"`
	// Participant's calculated role. (Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT)
	Role string `json:"role"`
	// Damage taken difference versus the calculated lane opponent(s) for a specified period.
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	// Damage taken for a specified period.
	DamageTakenPerMinDeltas map[string]float64 `json:"damageTakenPerMinDeltas"`
}

type LegacyMastery struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

type Matchlist struct {
	Matches    []MatchReference `json:"matches"`
	TotalGames int              `json:"totalGames"`
	StartIndex int              `json:"startIndex"`
	EndIndex   int              `json:"endIndex"`
}

type MatchReference struct {
	Lane       string `json:"lane"`
	GameID     int    `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Timestamp  int    `json:"timestamp"`
}
