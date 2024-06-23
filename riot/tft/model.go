package tft

// CurrentGameInfo contains current game information
type CurrentGameInfo struct {
	// The ID of the game
	GameID int64 `json:"gameId"`
	// The game type
	GameType string `json:"gameType"`
	// The game start time represented in epoch milliseconds
	GameStartTime int64 `json:"gameStartTime"`
	// The ID of the map
	MapID int64 `json:"mapId"`
	// The amount of time in seconds that has passed since the game started
	GameLength int64 `json:"gameLength"`
	// The ID of the platform on which the game is being played
	PlatformID string `json:"platformId"`
	// The game mode
	GameMode string `json:"gameMode"`
	// Banned champion information
	BannedChampions []BannedChampion `json:"bannedChamptions"`
	// The queue type (queue types are documented on the Game Constants page)
	GameQueueConfigID int64 `json:"gameQueueConfigId"`
	// The observer information
	Observers Observer `json:"observers"`
	// The participant information
	Participants []CurrentGameParticipant `json:"participants"`
}

// BannedChampion contains current game champion bans
type BannedChampion struct {
	// The turn during which the champion was banned
	PickTurn int `json:"pickTurn"`
	// The ID of the banned champion
	ChampionID int64 `json:"championId"`
	// The ID of the team that banned the champion
	TeamID int64 `json:"teamId"`
}

type Observer struct {
	// Key used to decrypt the spectator grid game data for playback
	EncryptionKey string `json:"encryptionKey"`
}

type CurrentGameParticipant struct {
	// The ID of the champion played by this participant
	ChampionID int64 `json:"championId"`
	// Perks/Runes Reforged Information
	Perks Perks `json:"perks"`
	// The ID of the profile icon used by this participant
	ProfileIconID int64 `json:"profileIconId"`
	// The team ID of this participant, indicating the participant's team
	TeamID int64 `json:"teamId"`
	// The encrypted summoner ID of this participant
	SummonerID string `json:"summonerId"`
	// The encrypted puuid of this participant
	PUUID string `json:"puuid"`
	// The ID of the first summoner spell used by this participant
	Spell1ID int64 `json:"spell1Id"`
	// The ID of the second summoner spell used by this participant
	Spell2ID int64 `json:"spell2Id"`
	// List of Game Customizations
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
}

type Perks struct {
	// IDs of the perks/runes assigned.
	PerkIDs []int64 `json:"perkIds"`
	// Primary runes path
	PerkStyle int64 `json:"perkStyle"`
	// Secondary runes path
	PerkSubStyle int64 `json:"perkSubStyle"`
}

type GameCustomizationObject struct {
	// Category identifier for Game Customization
	Category string `json:"category"`
	// Game Customization content
	Content string `json:"content"`
}

type FeaturedGames struct {
	// The list of featured games
	GameList []FeaturedGameInfo `json:"gameList"`
	// The suggested interval to wait before requesting FeaturedGames again
	ClientRefreshInterval int64 `json:"clientRefreshInterval"`
}

type FeaturedGameInfo struct {
	// The ID of the game
	GameID int64 `json:"gameId"`
	// The game type (Legal values: MATCHED)
	GameType string `json:"gameType"`
	// The ID of the map
	MapID int64 `json:"mapId"`
	// The amount of time in seconds that has passed since the game started
	GameLength int64 `json:"gameLength"`
	// The ID of the platform on which the game is being played
	PlatformID string `json:"platformId"`
	// The game mode (Legal values: TFT)
	GameMode string `json:"gameMode"`
	// Banned champion information
	BannedChampions []BannedChampion `json:"bannedChamptions"`
	// The queue type (queue types are documented on the Game Constants page)
	GameQueueConfigID int64 `json:"gameQueueConfigId"`
	// The observer information
	Observers Observer `json:"observers"`
	// The participant information
	Participants []FeaturedGameParticipant `json:"participants"`
}

type FeaturedGameParticipant struct {
	// The ID of the champion played by this participant
	ChampionID int64 `json:"championId"`
	// The ID of the profile icon used by this participant
	ProfileIconID int64 `json:"profileIconId"`
	// The team ID of this participant, indicating the participant's team
	TeamID int64 `json:"teamId"`
	// Encrypted summoner ID of this participant
	SummonerID string `json:"summonerId"`
	// Encrypted puuid of this participant
	PUUID string `json:"puuid"`
	// The ID of the first summoner spell used by this participant
	Spell1ID int64 `json:"spell1Id"`
	// The ID of the second summoner spell used by this participant
	Spell2ID int64 `json:"spell2Id"`
}

type LeagueList struct {
	LeagueID string       `json:"leagueId"`
	Entries  []LeagueItem `json:"entries"`
	Tier     string       `json:"tier"`
	Name     string       `json:"name"`
	Queue    string       `json:"queue"`
}

type LeagueItem struct {
	SummonerID   string      `json:"summonerId"`
	LeaguePoints int         `json:"leaguePoints"`
	Rank         string      `json:"rank"`
	Wins         int         `json:"wins"`
	Losses       int         `json:"losses"`
	Veteran      bool        `json:"veteran"`
	Inactive     bool        `json:"inactive"`
	FreshBlood   bool        `json:"freshBlood"`
	HotStreak    bool        `json:"hotStreak"`
	MiniSeries   []MiniSerie `json:"miniSeries"`
}

type MiniSerie struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type LeagueEntry struct {
	PUUID        string      `json:"puuid"`
	LeagueID     string      `json:"leagueId"`
	SummonerID   string      `json:"summonerId"`
	QueueType    string      `json:"queueType"`
	RatedTier    string      `json:"ratedTier"`
	RatedRating  int         `json:"ratedRating"`
	Tier         string      `json:"tier"`
	Rank         string      `json:"rank"`
	LeaguePoints int         `json:"leaguePoints"`
	Wins         int         `json:"wins"`
	Losses       int         `json:"losses"`
	HotStreak    bool        `json:"hotStreak"`
	Veteran      bool        `json:"veteran"`
	Inactive     bool        `json:"inactive"`
	FreshBlood   bool        `json:"freshBlood"`
	MiniSeries   []MiniSerie `json:"miniSeries"`
}

type TopRatedLadderEntry struct {
	SummonerID                   string `json:"summonerId"`
	RatedTier                    string `json:"ratedTier"`
	RatedRating                  int    `json:"ratedRating"`
	Wins                         int    `json:"wins"`
	PreviousUpdateLadderPosition int    `json:"previousUpdateLadderPosition"`
}

type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type Info struct {
	EndOfGameResult string        `json:"endOfGameResult"`
	GameCreation    int64         `json:"gameCreation"`
	GameID          int64         `json:"gameId"`
	GameDatetime    int64         `json:"game_datetime"`
	GameLength      float64       `json:"game_length"`
	GameVersion     string        `json:"game_version"`
	MapID           int64         `json:"mapId"`
	Participants    []Participant `json:"participants"`
	QueueID         int64         `json:"queueId"`
	InfoQueueID     int64         `json:"queue_id"`
	TFTGameType     string        `json:"tft_game_type"`
	TFTSetCoreName  string        `json:"tft_set_core_name"`
	TFTSetNumber    int64         `json:"tft_set_number"`
}

type Participant struct {
	Augments             []string         `json:"augments"`
	Companion            Companion        `json:"companion"`
	GoldLeft             int64            `json:"gold_left"`
	LastRound            int64            `json:"last_round"`
	Level                int64            `json:"level"`
	Missions             map[string]int64 `json:"missions"`
	Placement            int64            `json:"placement"`
	PlayersEliminated    int64            `json:"players_eliminated"`
	Puuid                string           `json:"puuid"`
	TimeEliminated       float64          `json:"time_eliminated"`
	TotalDamageToPlayers int64            `json:"total_damage_to_players"`
	Traits               []Trait          `json:"traits"`
	Units                []Unit           `json:"units"`
}

type Companion struct {
	ContentID string `json:"content_ID"`
	ItemID    int64  `json:"item_ID"`
	SkinID    int64  `json:"skin_ID"`
	Species   string `json:"species"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int64  `json:"num_units"`
	Style       int64  `json:"style"`
	TierCurrent int64  `json:"tier_current"`
	TierTotal   int64  `json:"tier_total"`
}

type Unit struct {
	CharacterID string   `json:"character_id"`
	ItemNames   []string `json:"itemNames"`
	Name        string   `json:"name"`
	Rarity      int64    `json:"rarity"`
	Tier        int64    `json:"tier"`
}

type Metadata struct {
	DataVersion  string   `json:"data_version"`
	MatchID      string   `json:"match_id"`
	Participants []string `json:"participants"`
}

type PlatformData struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Locales      []string         `json:"locales"`
	Maintenances []PlatformStatus `json:"maintenances"`
	Incidents    []PlatformStatus `json:"incidents"`
}

type PlatformStatus struct {
	ID                int               `json:"id"`
	MaintenanceStatus string            `json:"maintenanceStatus"`
	IncidentSeverity  string            `json:"incidentSeverity"`
	Titles            []PlatformContent `json:"titles"`
	Updates           []UpdateContent   `json:"updates"`
	CreatedAt         string            `json:"createdAt"`
	ArchiveAt         string            `json:"archiveAt"`
	UpdatedAt         string            `json:"updatedAt"`
	Platforms         []string          `json:"platforms"`
}

type PlatformContent struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type UpdateContent struct {
	ID               int               `json:"id"`
	Author           string            `json:"author"`
	Publish          bool              `json:"publish"`
	PublishLocations []string          `json:"publishLocations"`
	Translations     []PlatformContent `json:"translations"`
	CreatedAt        string            `json:"createdAt"`
	UpdatedAt        string            `json:"updatedAt"`
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	ProfileIconID int64  `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}
