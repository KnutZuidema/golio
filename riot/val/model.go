package val

// Act LocalizedNames is excluded because it is not sent when locale is set in request
type Act struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	IsActive bool   `json:"isActive"`
}

// ContentItem represents an individual content in ContentInfo
type ContentItem struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath"`
}

// ContentInfo represents all contents currently available in the game
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

// Content represents titles and translations fields of Update and Status
type Content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

// Update holds data of current software updates for the platforms
type Update struct {
	ID               int32      `json:"id"`
	Author           string     `json:"author"`
	Publish          bool       `json:"publish"`
	PublishLocations []string   `json:"publish_locations"`
	Translations     []*Content `json:"translations"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
}

// Status represents current maintenance and incidents
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

// PlatformData represents VALORANT status for given platform
type PlatformData struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Locales      []string  `json:"locales"`
	Maintenances []*Status `json:"maintenances"`
	Incidents    []*Status `json:"incidents"`
}

// Player holds data of individual players in a leaderboard
type Player struct {
	PuuID           string `json:"puuid"`
	GameName        string `json:"gameName"`
	TagLine         string `json:"tagLine"`
	LeaderboardRank int64  `json:"leaderboardRank"`
	RankedRating    int64  `json:"rankedRating"`
	NumberOfWins    int64  `json:"numberOfWins"`
}

// Leaderboard represents a leaderboard
type Leaderboard struct {
	Shard        string    `json:"shard"`
	ActID        string    `json:"actId"`
	TotalPlayers int64     `json:"totalPlayers"`
	Players      []*Player `json:"players"`
}

// Match represents a match with related data including players, match info and teams
type Match struct {
	MatchInfo    MatchInfo     `json:"matchInfo"`
	Players      []MatchPlayer `json:"players"`
	Coaches      []Coach       `json:"coaches"`
	Teams        []Team        `json:"teams"`
	RoundResults []RoundResult `json:"roundResults"`
}

// MatchInfo contains the data for a specific match
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

// MatchPlayer holds data of a player participating a match
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

// PlayerStats stats of a player in a match
type PlayerStats struct {
	Score          int          `json:"score"`
	RoundsPlayed   int          `json:"roundsPlayed"`
	Kills          int          `json:"kills"`
	Deaths         int          `json:"deaths"`
	Assists        int          `json:"assists"`
	PlaytimeMillis int          `json:"playtimeMillis"`
	AbilityCasts   AbilityCasts `json:"abilityCasts"`
}

// AbilityCasts number of casts of abilities for a player in a match
type AbilityCasts struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

// Coach holds coach id and team id
type Coach struct {
	PUUID  string `json:"puuid"`
	TeamID string `json:"teamId"`
}

// Team contains statistics of a team in a match
type Team struct {
	TeamID       string `json:"teamId"`
	Won          bool   `json:"won"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

// PlayerLocations represents player location for planting and defusing
type PlayerLocations struct {
	PUUID       string   `json:"puuid"`
	ViewRadians float32  `json:"viewRadians"`
	Location    Location `json:"location"`
}

// Location represents the location for PlayerLocations
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// RoundResult holds result data of a round
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

// PlayerRoundStats holds player stats by round
type PlayerRoundStats struct {
	PUUID   string   `json:"puuid"`
	Kills   []Kill   `json:"kills"`
	Damages []Damage `json:"damage"`
	Score   int      `json:"score"`
	Economy Economy  `json:"economy"`
	Ability Ability  `json:"ability"`
}

// Kill contains information of kills e.g. killer, victim, finishing damage
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

// FinishingDamage contains information of the finishing damage
type FinishingDamage struct {
	DamageType          string `json:"damageType"`
	DamageItem          string `json:"damageItem"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
}

// Damage contains information of a damage
type Damage struct {
	Receiver            string `json:"receiver"`
	Damage              bool   `json:"damage"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
	LegShots            bool   `json:"legshots"`
	BodyShots           bool   `json:"bodyshots"`
	Headshots           bool   `json:"headshots"`
}

// Economy holds economy information including spent credits
type Economy struct {
	LoadOutValue int    `json:"loadoutValue"`
	Weapon       string `json:"weapon"`
	Armor        string `json:"armor"`
	Remaining    int    `json:"remaining"`
	Spent        int    `json:"spent"`
}

// Ability holds ability effects of a player in a round
type Ability struct {
	GrenadeEffects  string `json:"grenadeEffects"`
	Ability1Effects string `json:"ability1Effects"`
	Ability2Effects string `json:"ability2Effects"`
	UltimateEffects string `json:"ultimateEffects"`
}

// MatchList represent match history of a player
type MatchList struct {
	PUUID   string           `json:"puuid"`
	History []MatchListEntry `json:"history"`
}

// MatchListEntry holds information regarding each match
type MatchListEntry struct {
	MatchID             string `json:"matchId"`
	GameStartTimeMillis int64  `json:"gameStartTimeMillis"`
	QueueID             string `json:"queueId"`
}

// RecentMatches represents last matches for live regions and e-sports routing
type RecentMatches struct {
	CurrentTime int64    `json:"currentTime"`
	MatchIDs    []string `json:"matchIds"`
}
