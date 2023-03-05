package val

// ActDto LocalizedNames is excluded because it is not sent when locale is set in request
type ActDto struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	IsActive bool   `json:"isActive"`
}

// ContentItemDto LocalizedNames is excluded because it is not sent when locale is set in request
type ContentItemDto struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	AssetName string `json:"assetName"`
	AssetPath string `json:"assetPath"`
}

type ContentInfoDto struct {
	Version      string            `json:"version"`
	Characters   []*ContentItemDto `json:"characters"`
	Maps         []*ContentItemDto `json:"maps"`
	Chromas      []*ContentItemDto `json:"chromas"`
	Skins        []*ContentItemDto `json:"skins"`
	SkinLevels   []*ContentItemDto `json:"skinLevels"`
	Equips       []*ContentItemDto `json:"equips"`
	GameModes    []*ContentItemDto `json:"gameModes"`
	Sprays       []*ContentItemDto `json:"sprays"`
	SprayLevels  []*ContentItemDto `json:"sprayLevels"`
	Charms       []*ContentItemDto `json:"charms"`
	CharmLevels  []*ContentItemDto `json:"charmLevels"`
	PlayerCards  []*ContentItemDto `json:"playerCards"`
	PlayerTitles []*ContentItemDto `json:"playerTitles"`
	Acts         []*ActDto         `json:"acts"`
}

type ContentDto struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type UpdateDto struct {
	Id               int32         `json:"id"`
	Author           string        `json:"author"`
	Publish          bool          `json:"publish"`
	PublishLocations []string      `json:"publish_locations"`
	Translations     []*ContentDto `json:"translations"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
}

type StatusDto struct {
	Id                int32         `json:"id"`
	MaintenanceStatus string        `json:"maintenance_status"`
	IncidentSeverity  string        `json:"incident_severity"`
	Titles            []*ContentDto `json:"titles"`
	Updates           []*UpdateDto  `json:"updates"`
	CreatedAt         string        `json:"created_at"`
	ArchiveAt         string        `json:"archive_at"`
	UpdatedAt         string        `json:"updated_at"`
	Platforms         []string      `json:"platforms"`
}

type PlatformDataDto struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Locales      []string     `json:"locales"`
	Maintenances []*StatusDto `json:"maintenances"`
	Incidents    []*StatusDto `json:"incidents"`
}

type PlayerDto struct {
	PuuId           string `json:"puuid"`
	GameName        string `json:"gameName"`
	TagLine         string `json:"tagLine"`
	LeaderboardRank int64  `json:"leaderboardRank"`
	RankedRating    int64  `json:"rankedRating"`
	NumberOfWins    int64  `json:"numberOfWins"`
}

type LeaderboardDto struct {
	Shard        string       `json:"shard"`
	ActId        string       `json:"actId"`
	TotalPlayers int64        `json:"totalPlayers"`
	Players      []*PlayerDto `json:"players"`
}
