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
	// Player's encrypted summonerId
	SummonerID   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
	Rank         string `json:"rank"`
	// First placement
	Wins int `json:"wins"`
	// Second through eighth placement.
	Losses     int         `json:"losses"`
	Veteran    bool        `json:"veteran"`
	Inactive   bool        `json:"inactive"`
	FreshBlood bool        `json:"freshBlood"`
	HotStreak  bool        `json:"hotStreak"`
	MiniSeries []MiniSerie `json:"miniSeries"`
}

type MiniSerie struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type LeagueEntry struct {
	// Player Universal Unique Identifier. Exact length of 78 characters. (Encrypted)
	PUUID string `json:"puuid"`
	// Not included for the RANKED_TFT_TURBO queueType.
	LeagueID string `json:"leagueId"`
	// Player's encrypted summonerId
	SummonerID string `json:"summonerId"`
	QueueType  string `json:"queueType"`
	// Only included for the RANKED_TFT_TURBO queueType. (Legal values: ORANGE, PURPLE, BLUE, GREEN, GRAY)
	RatedTier string `json:"ratedTier"`
	// Only included for the RANKED_TFT_TURBO queueType.
	RatedRating int `json:"ratedRating"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Tier string `json:"tier"`
	// The player's division within a tier. Not included for the RANKED_TFT_TURBO queueType.
	Rank string `json:"rank"`
	// Not included for the RANKED_TFT_TURBO queueType.
	LeaguePoints int `json:"leaguePoints"`
	// First placement
	Wins int `json:"wins"`
	// Second through eighth placement
	Losses int `json:"losses"`
	// Not included for the RANKED_TFT_TURBO queueType.
	HotStreak bool `json:"hotStreak"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Veteran bool `json:"veteran"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Inactive bool `json:"inactive"`
	// Not included for the RANKED_TFT_TURBO queueType.
	FreshBlood bool `json:"freshBlood"`
	// Not included for the RANKED_TFT_TURBO queueType.
	MiniSeries []MiniSerie `json:"miniSeries"`
}

type TopRatedLadderEntry struct {
	// Player's encrypted summonerId.
	SummonerID string `json:"summonerId"`
	// (Legal values: ORANGE, PURPLE, BLUE, GREEN, GRAY)
	RatedTier   string `json:"ratedTier"`
	RatedRating int    `json:"ratedRating"`
	// First placement
	Wins                         int `json:"wins"`
	PreviousUpdateLadderPosition int `json:"previousUpdateLadderPosition"`
}

type Match struct {
	// Match metadata
	Metadata Metadata `json:"metadata"`
	// Match info
	Info MatchInfo `json:"info"`
}

type MatchInfo struct {
	// Unix timestamp
	GameDatetime int64 `json:"game_datetime"`
	// Game length in seconds
	GameLength float64 `json:"game_length"`
	// Game variation key. Game variations documented in TFT static data
	GameVariation string `json:"game_variation"`
	// Game client version
	GameVersion  string        `json:"game_version"`
	Participants []Participant `json:"participants"`
	// Please refer to the League of Legends documentation
	QueueID int64 `json:"queueId"`
	// Teamfight Tactics set number
	TFTSetNumber int64 `json:"tft_set_number"`
}

type Participant struct {
	// Participant's companion
	Companion Companion `json:"companion"`
	// Gold left after participant was eliminated
	GoldLeft int64 `json:"gold_left"`
	// The round the participant was eliminated in. Note: If the player was eliminated in stage 2-1 their last_round would be 5.
	LastRound int64 `json:"last_round"`
	// Participant Little Legend level. Note: This is not the number of active units
	Level int64 `json:"level"`
	// Participant placement upon elimination
	Placement int64 `json:"placement"`
	// Number of players the participant eliminated.
	PlayersEliminated int64  `json:"players_eliminated"`
	PUUID             string `json:"puuid"`
	// The number of seconds before the participant was eliminated
	TimeEliminated float64 `json:"time_eliminated"`
	// Damage the participant dealt to other players.
	TotalDamageToPlayers int64 `json:"total_damage_to_players"`
	// A complete list of traits for the participant's active units
	Traits []Trait `json:"traits"`
	// A list of active units for the participant
	Units []Unit `json:"units"`
}

type Companion struct {
	ContentID string `json:"content_ID"`
	ItemID    int64  `json:"item_ID"`
	SkinID    int64  `json:"skin_ID"`
	Species   string `json:"species"`
}

type Trait struct {
	// Trait name
	Name string `json:"name"`
	// Number of units with this trait
	NumUnits int64 `json:"num_units"`
	// Current style for this trait. (0 = No style, 1 = Bronze, 2 = Silver, 3 = Gold, 4 = Chromatic)
	Style int64 `json:"style"`
	// Current active tier for the trait
	TierCurrent int64 `json:"tier_current"`
	// Total tiers for the trait
	TierTotal int64 `json:"tier_total"`
}

type Unit struct {
	// This field was introduced in patch 9.22 with data_version 2.
	CharacterID string `json:"character_id"`
	// A list of the unit's items. Please refer to the Teamfight Tactics documentation for item ids.
	Items []int `json:"itemNames"`
	// If a unit is chosen as part of the Fates set mechanic, the chosen trait will be indicated by this field. Otherwise this field is excluded from the response.
	Chosen string `json:"chosen"`
	Name   string `json:"name"`
	// Unit rarity. This doesn't equate to the unit cost.
	Rarity int `json:"rarity"`
	// Unit tier
	Tier int `json:"tier"`
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
	ID int `json:"id"`
	// (Legal values: scheduled, in_progress, complete)
	MaintenanceStatus string `json:"maintenanceStatus"`
	// (Legal values: info, warning, critical)
	IncidentSeverity string            `json:"incidentSeverity"`
	Titles           []PlatformContent `json:"titles"`
	Updates          []UpdateContent   `json:"updates"`
	CreatedAt        string            `json:"createdAt"`
	ArchiveAt        string            `json:"archiveAt"`
	UpdatedAt        string            `json:"updatedAt"`
	//	(Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms []string `json:"platforms"`
}

type PlatformContent struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type UpdateContent struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Publish bool   `json:"publish"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string          `json:"publishLocations"`
	Translations     []PlatformContent `json:"translations"`
	CreatedAt        string            `json:"createdAt"`
	UpdatedAt        string            `json:"updatedAt"`
}

type Summoner struct {
	//Encrypted summoner ID. Max length 63 characters.
	ID string `json:"id"`
	// Encrypted account ID. Max length 56 characters.
	AccountID string `json:"accountId"`
	// Encrypted PUUID. Exact length of 78 characters.
	PUUID string `json:"puuid"`
	// ID of the summoner icon associated with the summoner.
	ProfileIconID int64 `json:"profileIconId"`
	// Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: summoner name change, summoner level change, or profile icon change.
	RevisionDate int64 `json:"revisionDate"`
	// Summoner level associated with the summoner.
	SummonerLevel int64 `json:"summonerLevel"`
}
