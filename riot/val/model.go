package val

// Act LocalizedNames is excluded because it is not sent when locale is set in request
type Act struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	IsActive bool   `json:"isActive"`
}

// ContentItem LocalizedNames is excluded because it is not sent when locale is set in request
type ContentItem struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath"`
}

type ContentInfo struct {
	Version      string         `json:"version"`
	Characters   []*ContentItem `json:"characters"`
	Maps         []*ContentItem `json:"maps"`
	Chromas      []*ContentItem `json:"chromas"`
	Skins        []*ContentItem `json:"skins"`
	SkinLevels   []*ContentItem `json:"skinLevels"`
	Equips       []*ContentItem `json:"equips"`
	GameModes    []*ContentItem `json:"gameModes"`
	Sprays       []*ContentItem `json:"sprays"`
	SprayLevels  []*ContentItem `json:"sprayLevels"`
	Charms       []*ContentItem `json:"charms"`
	CharmLevels  []*ContentItem `json:"charmLevels"`
	PlayerCards  []*ContentItem `json:"playerCards"`
	PlayerTitles []*ContentItem `json:"playerTitles"`
	Acts         []*Act         `json:"acts"`
}

type Content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type Update struct {
	ID               int32      `json:"id"`
	Author           string     `json:"author"`
	Publish          bool       `json:"publish"`
	PublishLocations []string   `json:"publish_locations"`
	Translations     []*Content `json:"translations"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
}

type Status struct {
	ID                int32      `json:"id"`
	MaintenanceStatus string     `json:"maintenance_status"`
	IncidentSeverity  string     `json:"incident_severity"`
	Titles            []*Content `json:"titles"`
	Updates           []*Update  `json:"updates"`
	CreatedAt         string     `json:"created_at"`
	ArchiveAt         string     `json:"archive_at"`
	UpdatedAt         string     `json:"updated_at"`
	Platforms         []string   `json:"platforms"`
}

type PlatformData struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Locales      []string  `json:"locales"`
	Maintenances []*Status `json:"maintenances"`
	Incidents    []*Status `json:"incidents"`
}

type Player struct {
	PuuID           string `json:"puuid"`
	GameName        string `json:"gameName"`
	TagLine         string `json:"tagLine"`
	LeaderboardRank int64  `json:"leaderboardRank"`
	RankedRating    int64  `json:"rankedRating"`
	NumberOfWins    int64  `json:"numberOfWins"`
}

type Leaderboard struct {
	Shard        string    `json:"shard"`
	ActID        string    `json:"actId"`
	TotalPlayers int64     `json:"totalPlayers"`
	Players      []*Player `json:"players"`
}

type Match struct {
	MatchInfo    MatchInfo     `json:"matchInfo"`
	Players      []MatchPlayer `json:"players"`
	Coaches      []Coach       `json:"coaches"`
	Teams        []Team        `json:"teams"`
	RoundResults []RoundResult `json:"roundResults"`
}

type MatchInfo struct {
	MatchID            string `json:"matchId"`
	MapID              string `json:"mapId"`
	GameLengthMillis   int    `json:"gameLengthMillis"`
	GameStartMillis    int64  `json:"gameStartMillis"`
	ProvisioningFlowID string `json:"provisioningFlowId"`
	IsCompleted        bool   `json:"isCompleted"`
	CustomGameName     string `json:"customGameName"`
	QueueID            string `json:"queueId"`
	GameMode           string `json:"gameMode"`
	IsRanked           bool   `json:"isRanked"`
	SeasonID           string `json:"seasonId"`
}

type MatchPlayer struct {
	PuuID           string      `json:"puuid"`
	GameName        string      `json:"gameName"`
	TagLine         string      `json:"tagLine"`
	TeamID          string      `json:"teamId"`
	PartyID         string      `json:"partyId"`
	CharacterID     string      `json:"characterId"`
	Stats           PlayerStats `json:"stats"`
	CompetitiveTier int         `json:"competitiveTier"`
	PlayerCard      string      `json:"playerCard"`
	PlayerTitle     string      `json:"playerTitle"`
}

type PlayerStats struct {
	Score          int          `json:"score"`
	RoundsPlayed   int          `json:"roundsPlayed"`
	Kills          int          `json:"kills"`
	Deaths         int          `json:"deaths"`
	Assists        int          `json:"assists"`
	PlaytimeMillis int          `json:"playtimeMillis"`
	AbilityCasts   AbilityCasts `json:"abilityCasts"`
}

type AbilityCasts struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

type Coach struct {
	PUUID  string `json:"puuid"`
	TeamID string `json:"teamId"`
}

type Team struct {
	TeamID       string `json:"teamId"`
	Won          bool   `json:"won"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

type PlayerLocations struct {
	PUUID       string   `json:"puuid"`
	ViewRadians float32  `json:"viewRadians"`
	Location    Location `json:"location"`
}

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type RoundResult struct {
	RoundNum              int                `json:"roundNum"`
	RoundResult           string             `json:"roundResult"`
	RoundCeremony         string             `json:"roundCeremony"`
	WinningTeam           string             `json:"winningTeam"`
	BombPlanter           string             `json:"bombPlanter"`
	BombDefuser           string             `json:"bombDefuser"`
	PlantRoundTime        int                `json:"plantRoundTime"`
	PlantPlayerLocations  []PlayerLocations  `json:"plantPlayerLocations"`
	PlantLocation         Location           `json:"plantLocation"`
	PlantSite             string             `json:"plantSite"`
	DefuseRoundTime       int                `json:"defuseRoundTime"`
	DefusePlayerLocations []PlayerLocations  `json:"defusePlayerLocations"`
	DefuseLocation        Location           `json:"defuseLocation"`
	PlayerStats           []PlayerRoundStats `json:"playerStats"`
	RoundResultCode       string             `json:"roundResultCode"`
}

type PlayerRoundStats struct {
	PUUID   string   `json:"puuid"`
	Kills   []Kill   `json:"kills"`
	Damages []Damage `json:"damage"`
	Score   int      `json:"score"`
	Economy Economy  `json:"economy"`
	Ability Ability  `json:"ability"`
}

type Kill struct {
	TimeSinceGameStartMillis  int               `json:"timeSinceGameStartMillis"`
	TimeSinceRoundStartMillis int               `json:"timeSinceRoundStartMillis"`
	Killer                    string            `json:"killer"`
	Victim                    string            `json:"victim"`
	VictimLocation            Location          `json:"victimLocation"`
	Assistants                []string          `json:"assistants"`
	PlayerLocations           []PlayerLocations `json:"playerLocations"`
	FinishingDamage           FinishingDamage   `json:"finishingDamage"`
}

type FinishingDamage struct {
	DamageType          string `json:"damageType"`
	DamageItem          string `json:"damageItem"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
}

type Damage struct {
	Receiver            string `json:"receiver"`
	Damage              bool   `json:"damage"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
	LegShots            bool   `json:"legshots"`
	BodyShots           bool   `json:"bodyshots"`
	Headshots           bool   `json:"headshots"`
}

type Economy struct {
	LoadOutValue int    `json:"loadoutValue"`
	Weapon       string `json:"weapon"`
	Armor        string `json:"armor"`
	Remaining    int    `json:"remaining"`
	Spent        int    `json:"spent"`
}

type Ability struct {
	GrenadeEffects  string `json:"grenadeEffects"`
	Ability1Effects string `json:"ability1Effects"`
	Ability2Effects string `json:"ability2Effects"`
	UltimateEffects string `json:"ultimateEffects"`
}

type MatchList struct {
	PUUID           string           `json:"puuid"`
	Ability1Effects []MatchListEntry `json:"history"`
}

type MatchListEntry struct {
	MatchId             string `json:"matchId"`
	GameStartTimeMillis int64  `json:"gameStartTimeMillis"`
	QueueId             string `json:"queueId"`
}

type RecentMatches struct {
	CurrentTime int64    `json:"currentTime"`
	MatchIDs    []string `json:"matchIds"`
}
