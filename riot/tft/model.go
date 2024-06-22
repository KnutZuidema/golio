package tft

type CurrentGameInfo struct {
	GameID            int64                    `json:"gameId"`
	GameType          string                   `json:"gameType"`
	GameStartTime     int64                    `json:"gameStartTime"`
	MapID             int64                    `json:"mapId"`
	GameLength        int64                    `json:"gameLength"`
	PlatformID        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	BannedChampions   []BannedChampion         `json:"bannedChamptions"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
}

type BannedChampion struct {
	PickTurn   int   `json:"pickTurn"`
	ChampionID int64 `json:"championId"`
	TeamID     int64 `json:"teamId"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

type CurrentGameParticipant struct {
	ChampionID               int64                     `json:"championId"`
	Perks                    Perks                     `json:"perks"`
	ProfileIconID            int64                     `json:"profileIconId"`
	TeamID                   int64                     `json:"teamId"`
	SummonerID               string                    `json:"summonerId"`
	PUUID                    string                    `json:"puuid"`
	Spell1ID                 int64                     `json:"spell1Id"`
	Spell2ID                 int64                     `json:"spell2Id"`
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
}

type Perks struct {
	PerkIDs      []int64 `json:"perkIds"`
	PerkStyle    int64   `json:"perkStyle"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

type FeaturedGames struct {
	GameList              []FeaturedGameInfo `json:"gameList"`
	ClientRefreshInterval int64              `json:"clientRefreshInterval"`
}

type FeaturedGameInfo struct {
	GameID            int64                     `json:"gameId"`
	GameType          string                    `json:"gameType"`
	MapID             int64                     `json:"mapId"`
	GameLength        int64                     `json:"gameLength"`
	PlatformID        string                    `json:"platformId"`
	GameMode          string                    `json:"gameMode"`
	BannedChampions   []BannedChampion          `json:"bannedChamptions"`
	GameQueueConfigID int64                     `json:"gameQueueConfigId"`
	Observers         Observer                  `json:"observers"`
	Participants      []FeaturedGameParticipant `json:"participants"`
}

type FeaturedGameParticipant struct {
	ChampionID    int64  `json:"championId"`
	ProfileIconID int64  `json:"profileIconId"`
	TeamID        int64  `json:"teamId"`
	SummonerID    string `json:"summonerId"`
	PUUID         string `json:"puuid"`
	Spell1ID      int64  `json:"spell1Id"`
	Spell2ID      int64  `json:"spell2Id"`
}
