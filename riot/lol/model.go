package lol

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/KnutZuidema/golio/datadragon"
	"github.com/KnutZuidema/golio/static"
)

// ChampionInfo contains information about the free champion rotation
type ChampionInfo struct {
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIDsForNewPlayers"`
	FreeChampionIDs              []int `json:"freeChampionIDs"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// GetChampionsForNewPlayers returns data for champions available for free to new players
func (i *ChampionInfo) GetChampionsForNewPlayers(client *datadragon.Client) ([]datadragon.ChampionDataExtended, error) {
	res := make([]datadragon.ChampionDataExtended, 0, len(i.FreeChampionIDsForNewPlayers))
	for _, id := range i.FreeChampionIDsForNewPlayers {
		champion, err := client.GetChampionByID(strconv.Itoa(id))
		if err != nil {
			return nil, err
		}
		res = append(res, champion)
	}
	return res, nil
}

// GetChampions returns data for champions available for free
func (i *ChampionInfo) GetChampions(client *datadragon.Client) ([]datadragon.ChampionDataExtended, error) {
	res := make([]datadragon.ChampionDataExtended, 0, len(i.FreeChampionIDsForNewPlayers))
	for _, id := range i.FreeChampionIDs {
		champion, err := client.GetChampionByID(strconv.Itoa(id))
		if err != nil {
			return nil, err
		}
		res = append(res, champion)
	}
	return res, nil
}

// ChampionMastery represents the mastery of a champion in the mastery system for a summoner
type ChampionMastery struct {
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionID                   int    `json:"championId"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	LastPlayTime                 int    `json:"lastPlayTime"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	SummonerID                   string `json:"summonerId"`
}

// GetSummoner returns the summoner of this mastery
func (m *ChampionMastery) GetSummoner(client *Client) (*Summoner, error) {
	return client.Summoner.GetByID(m.SummonerID)
}

// GetChampion returns the champion of this mastery
func (m *ChampionMastery) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(m.ChampionID))
}

// LeagueList represents a league containing all player entries in it
type LeagueList struct {
	LeagueID      string        `json:"leagueId"`
	Tier          string        `json:"tier"`
	Entries       []*LeagueItem `json:"entries"`
	Queue         string        `json:"queue"`
	Name          string        `json:"name"`
	sortedEntries []*LeagueItem
}

// GetRank returns the entry at the given rank, sorted by league points
func (l *LeagueList) GetRank(i int) *LeagueItem {
	if l.sortedEntries == nil || len(l.sortedEntries) != len(l.Entries) {
		l.sortedEntries = make([]*LeagueItem, len(l.Entries))
		copy(l.sortedEntries, l.Entries)
		sort.Slice(
			l.sortedEntries, func(i, j int) bool {
				return l.sortedEntries[i].LeaguePoints > l.sortedEntries[j].LeaguePoints
			},
		)
	}
	return l.sortedEntries[i]
}

// LeagueItem represents a summoners ranked position in a league
type LeagueItem struct {
	QueueType    string      `json:"queueType"`
	SummonerName string      `json:"summonerName"`
	HotStreak    bool        `json:"hotStreak"`
	MiniSeries   *MiniSeries `json:"miniSeries"`
	Wins         int         `json:"wins"`
	Veteran      bool        `json:"veteran"`
	Losses       int         `json:"losses"`
	FreshBlood   bool        `json:"freshBlood"`
	Inactive     bool        `json:"inactive"`
	Tier         string      `json:"tier"`
	Rank         string      `json:"rank"`
	SummonerID   string      `json:"summonerId"`
	LeaguePoints int         `json:"leaguePoints"`
}

// GetSummoner returns the summoner of this league item
func (i *LeagueItem) GetSummoner(client *Client) (*Summoner, error) {
	return client.Summoner.GetByID(i.SummonerID)
}

// MiniSeries represents a mini series when playing to ascend to the next ranked tier
type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

// Match contains information about a match
type Match struct {
	// Match metadata
	Metadata *MatchMetadata `json:"metadata"`
	// Match info
	Info *MatchInfo `json:"info"`
}

// MatchMetadata contains metadata for a specific match
type MatchMetadata struct {
	// Match data version
	DataVersion string `json:"dataVersion"`
	// Match ID
	MatchID string `json:"matchId"`
	// List of participant PUUIDs
	Participants []string `json:"participants"`
}

// MatchInfo contains the data for a specific match
type MatchInfo struct {
	// Unix timestamp for when the game is created on the game server (i.e., the loading screen).
	GameCreation int64 `json:"gameCreation"`
	// Prior to patch 11.20, this field returns the game length in milliseconds calculated
	// from gameEndTimestamp - gameStartTimestamp. Post patch 11.20, this field returns the max
	// timePlayed of any participant in the game in seconds, which makes the behavior of this
	// field consistent with that of match-v4. The best way to handling the change in this field
	// is to treat the value as milliseconds if the gameEndTimestamp field isn't in the response
	// and to treat the value as seconds if gameEndTimestamp is in the response.
	GameDuration int `json:"gameDuration"`
	// Unix timestamp for when match ends on the game server. This timestamp can occasionally
	// be significantly longer than when the match "ends". The most reliable way of determining
	// the timestamp for the end of the match would be to add the max time played of any
	// participant to the gameStartTimestamp. This field was added to match-v5 in patch 11.20 on Oct 5th, 2021.
	GameEndTimestamp int64 `json:"gameEndTimestamp"`
	GameID           int64 `json:"gameId"`
	// Please refer to the Game Constants documentation.
	GameMode string `json:"gameMode"`
	GameName string `json:"gameName"`
	// Unix timestamp for when match starts on the game server.
	GameStartTimestamp int64 `json:"gameStartTimestamp"`
	// Please refer to the Game Constants documentation.
	GameType string `json:"gameType"`
	// The first two parts can be used to determine the patch a game was played on.
	GameVersion string `json:"gameVersion"`
	// Please refer to the Game Constants documentation.
	MapID int `json:"mapId"`
	// Participant information.
	Participants []*Participant `json:"participants"`
	// Platform where the match was played.
	PlatformID string `json:"platformId"`
	// Please refer to the Game Constants documentation.
	QueueID int `json:"queueId"`
	// Team information.
	Teams []*Team `json:"teams"`
	// Tournament code used to generate the match. This field was added to match-v5 in patch 11.13 on June 23rd, 2021.
	TournamentCode string `json:"tournamentCode"`
}

// GetQueue returns the queue this match was played in
func (m *MatchInfo) GetQueue(client *static.Client) (static.Queue, error) {
	return client.GetQueue(m.QueueID)
}

// GetMap returns the map this match was played on
func (m *MatchInfo) GetMap(client *static.Client) (static.Map, error) {
	return client.GetMap(m.MapID)
}

// GetGameType returns the gameType this match was played in
func (m *MatchInfo) GetGameType(client *static.Client) (static.GameType, error) {
	return client.GetGameType(m.GameType)
}

// GetGameMode returns the gameMode this match was played in
func (m *MatchInfo) GetGameMode(client *static.Client) (static.GameMode, error) {
	return client.GetGameMode(m.GameMode)
}

// StatPerks hold stats for a perk
type StatPerks struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

// Selections contains information about perk selections
type Selections struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

// Styles holds perk style information
type Styles struct {
	Description string       `json:"description"`
	Selections  []Selections `json:"selections"`
	Style       int          `json:"style"`
}

// ParticipantPerks holds the perks for a participant in a match
type ParticipantPerks struct {
	StatPerks *StatPerks `json:"statPerks"`
	Styles    []Styles   `json:"styles"`
}

// Participant hold information for a participant of a match
type Participant struct {
	Assists         int `json:"assists"`
	BaronKills      int `json:"baronKills"`
	BountyLevel     int `json:"bountyLevel"`
	ChampExperience int `json:"champExperience"`
	ChampLevel      int `json:"champLevel"`
	// Prior to patch 11.4, on Feb 18th, 2021, this field returned invalid championIds.
	// We recommend determining the champion based on the championName field for matches played prior to patch 11.4.
	ChampionID   int    `json:"championId"`
	ChampionName string `json:"championName"`
	// This field is currently only utilized for Kayn's transformations.
	// (Legal values: 0 - None, 1 - Slayer, 2 - Assassin)
	ChampionTransform         int  `json:"championTransform"`
	ConsumablesPurchased      int  `json:"consumablesPurchased"`
	DamageDealtToBuildings    int  `json:"damageDealtToBuildings"`
	DamageDealtToObjectives   int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets      int  `json:"damageDealtToTurrets"`
	DamageSelfMitigated       int  `json:"damageSelfMitigated"`
	Deaths                    int  `json:"deaths"`
	DetectorWardsPlaced       int  `json:"detectorWardsPlaced"`
	DoubleKills               int  `json:"doubleKills"`
	DragonKills               int  `json:"dragonKills"`
	FirstBloodAssist          bool `json:"firstBloodAssist"`
	FirstBloodKill            bool `json:"firstBloodKill"`
	FirstTowerAssist          bool `json:"firstTowerAssist"`
	FirstTowerKill            bool `json:"firstTowerKill"`
	GameEndedInEarlySurrender bool `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender      bool `json:"gameEndedInSurrender"`
	GoldEarned                int  `json:"goldEarned"`
	GoldSpent                 int  `json:"goldSpent"`
	// Both individualPosition and teamPosition are computed by the game server and are
	// different versions of the most likely position played by a player. The individualPosition
	// is the best guess for which position the player actually played in isolation of
	// anything else. The teamPosition is the best guess for which position the player
	// actually played if we add the constraint that each team must have one top player, one
	// jungle, one middle, etc. Generally the recommendation is to use the teamPosition field
	// over the individualPosition field.
	IndividualPosition             string            `json:"individualPosition"`
	InhibitorKills                 int               `json:"inhibitorKills"`
	InhibitorTakedowns             int               `json:"inhibitorTakedowns"`
	InhibitorsLost                 int               `json:"inhibitorsLost"`
	Item0                          int               `json:"item0"`
	Item1                          int               `json:"item1"`
	Item2                          int               `json:"item2"`
	Item3                          int               `json:"item3"`
	Item4                          int               `json:"item4"`
	Item5                          int               `json:"item5"`
	Item6                          int               `json:"item6"`
	ItemsPurchased                 int               `json:"itemsPurchased"`
	KillingSprees                  int               `json:"killingSprees"`
	Kills                          int               `json:"kills"`
	Lane                           string            `json:"lane"`
	LargestCriticalStrike          int               `json:"largestCriticalStrike"`
	LargestKillingSpree            int               `json:"largestKillingSpree"`
	LargestMultiKill               int               `json:"largestMultiKill"`
	LongestTimeSpentLiving         int               `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int               `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int               `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int               `json:"magicDamageTaken"`
	NeutralMinionsKilled           int               `json:"neutralMinionsKilled"`
	NexusKills                     int               `json:"nexusKills"`
	NexusLost                      int               `json:"nexusLost"`
	NexusTakedowns                 int               `json:"nexusTakedowns"`
	ObjectivesStolen               int               `json:"objectivesStolen"`
	ObjectivesStolenAssists        int               `json:"objectivesStolenAssists"`
	ParticipantID                  int               `json:"participantId"`
	PentaKills                     int               `json:"pentaKills"`
	Perks                          *ParticipantPerks `json:"perks"`
	PhysicalDamageDealt            int               `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int               `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int               `json:"physicalDamageTaken"`
	ProfileIcon                    int               `json:"profileIcon"`
	PUUID                          string            `json:"puuid"`
	QuadraKills                    int               `json:"quadraKills"`
	RiotIDGameName                 string            `json:"riotIdGameName"`
	RiotIDName                     string            `json:"riotIdName"`
	RiotIDTagline                  string            `json:"riotIdTagline"`
	Role                           string            `json:"role"`
	SightWardsBoughtInGame         int               `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int               `json:"spell1Casts"`
	Spell2Casts                    int               `json:"spell2Casts"`
	Spell3Casts                    int               `json:"spell3Casts"`
	Spell4Casts                    int               `json:"spell4Casts"`
	Summoner1Casts                 int               `json:"summoner1Casts"`
	Summoner1ID                    int               `json:"summoner1Id"`
	Summoner2Casts                 int               `json:"summoner2Casts"`
	Summoner2ID                    int               `json:"summoner2Id"`
	SummonerID                     string            `json:"summonerId"`
	SummonerLevel                  int               `json:"summonerLevel"`
	SummonerName                   string            `json:"summonerName"`
	TeamEarlySurrendered           bool              `json:"teamEarlySurrendered"`
	TeamID                         int               `json:"teamId"`
	// Both individualPosition and teamPosition are computed by the game server and are
	// different versions of the most likely position played by a player. The individualPosition
	// is the best guess for which position the player actually played in isolation of
	// anything else. The teamPosition is the best guess for which position the player
	// actually played if we add the constraint that each team must have one top player, one
	// jungle, one middle, etc. Generally the recommendation is to use the teamPosition field
	// over the individualPosition field.
	TeamPosition                   string `json:"teamPosition"`
	TimeCCingOthers                int    `json:"timeCCingOthers"`
	TimePlayed                     int    `json:"timePlayed"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TripleKills                    int    `json:"tripleKills"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	TurretKills                    int    `json:"turretKills"`
	TurretTakedowns                int    `json:"turretTakedowns"`
	TurretsLost                    int    `json:"turretsLost"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionScore                    int    `json:"visionScore"`
	VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
	WardsKilled                    int    `json:"wardsKilled"`
	WardsPlaced                    int    `json:"wardsPlaced"`
	Win                            bool   `json:"win"`
}

// GetSummoner returns the summoner info for this player
func (p *Participant) GetSummoner(client *Client) (*Summoner, error) {
	return client.Summoner.GetByPUUID(p.PUUID)
}

// GetProfileIcon returns the profile icon data for this player
func (p *Participant) GetProfileIcon(client *datadragon.Client) (datadragon.ProfileIcon, error) {
	return client.GetProfileIcon(p.ProfileIcon)
}

// GetChampion returns the champion played by this participant
func (p *Participant) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(p.ChampionID))
}

// GetSpell1 returns the first summoner spell of this participant
func (p *Participant) GetSpell1(client *datadragon.Client) (datadragon.SummonerSpell, error) {
	return client.GetSummonerSpell(strconv.Itoa(p.Summoner1ID))
}

// GetSpell2 returns the second summoner spell of this participant
func (p *Participant) GetSpell2(client *datadragon.Client) (datadragon.SummonerSpell, error) {
	return client.GetSummonerSpell(strconv.Itoa(p.Summoner2ID))
}

// GetItem0 returns the item in slot 0
func (p *Participant) GetItem0(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item0))
}

// GetItem1 returns the item in slot 1
func (p *Participant) GetItem1(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item1))
}

// GetItem2 returns the item in slot 2
func (p *Participant) GetItem2(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item2))
}

// GetItem3 returns the item in slot 3
func (p *Participant) GetItem3(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item3))
}

// GetItem4 returns the item in slot 4
func (p *Participant) GetItem4(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item4))
}

// GetItem5 returns the item in slot 5
func (p *Participant) GetItem5(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item5))
}

// GetItem6 returns the item in slot 6
func (p *Participant) GetItem6(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(p.Item6))
}

// TeamBan is a champion banned by a team
type TeamBan struct {
	// Turn during which the champion was banned.
	PickTurn int `json:"pickTurn"`
	// Banned championId.
	ChampionID int `json:"championId"`
}

// GetChampion returns the champion that was banned
func (b *TeamBan) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(b.ChampionID))
}

// Objective holds information for a single objective
type Objective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

// Objectives holds info for a teeam's objeectives
type Objectives struct {
	Baron      Objective `json:"baron"`
	Champion   Objective `json:"champion"`
	Dragon     Objective `json:"dragon"`
	Inhibitor  Objective `json:"inhibitor"`
	RiftHerald Objective `json:"riftHerald"`
	Tower      Objective `json:"tower"`
}

// Team holds information for a team in a match
type Team struct {
	Bans       []*TeamBan `json:"bans"`
	Objectives Objectives `json:"objectives"`
	TeamID     int        `json:"teamId"`
	Win        bool       `json:"win"`
}

// MatchTimeline contains timeline frames for a match
type MatchTimeline struct {
	Frames   []*MatchFrame `json:"frames"`
	Interval int           `json:"frameInterval"`
}

// MatchFrame is a single frame in the timeline of a game
type MatchFrame struct {
	Timestamp         int                          `json:"timestamp"`
	ParticipantFrames map[string]*ParticipantFrame `json:"participantFrames"`
	Events            []*MatchEvent                `json:"events"`
}

// ParticipantFrame contains information about a participant in a game at a single timestamp
type ParticipantFrame struct {
	TotalGold           int            `json:"totalGold"`
	TeamScore           int            `json:"teamScore"`
	ParticipantID       int            `json:"participantId"`
	Level               int            `json:"level"`
	CurrentGold         int            `json:"currentGold"`
	MinionsKilled       int            `json:"minionsKilled"`
	DominionScore       int            `json:"dominionScore"`
	Position            *MatchPosition `json:"position"`
	XP                  int            `json:"xp"`
	JungleMinionsKilled int            `json:"jungleMinionsKilled"`
}

// MatchEventType is the type of an event
type MatchEventType string

// All legal value for match event types
const (
	MatchEventTypeChampionKill     MatchEventType = "CHAMPION_KILL"
	MatchEventTypeWardPlaced                      = "WARD_PLACED"
	MatchEventTypeWardKill                        = "WARD_KILL"
	MatchEventTypeBuildingKill                    = "BUILDING_KILL"
	MatchEventTypeEliteMonsterKill                = "ELITE_MONSTER_KILL"
	MatchEventTypeItemPurchased                   = "ITEM_PURCHASED"
	MatchEventTypeItemSold                        = "ITEM_SOLD"
	MatchEventTypeItemDestroyed                   = "ITEM_DESTROYED"
	MatchEventTypeItemUndo                        = "ITEM_UNDO"
	MatchEventTypeSkillLevelUp                    = "SKILL_LEVEL_UP"
	MatchEventTypeAscendedEvent                   = "ASCENDED_EVENT"
	MatchEventTypeCapturePoint                    = "CAPTURE_POINT"
	MatchEventTypePoroKingSummon                  = "PORO_KING_SUMMON"
)

var (
	// MatchEventTypes is a list of all available match events
	MatchEventTypes = []MatchEventType{
		MatchEventTypeChampionKill,
		MatchEventTypeWardPlaced,
		MatchEventTypeWardKill,
		MatchEventTypeBuildingKill,
		MatchEventTypeEliteMonsterKill,
		MatchEventTypeItemPurchased,
		MatchEventTypeItemSold,
		MatchEventTypeItemDestroyed,
		MatchEventTypeItemUndo,
		MatchEventTypeSkillLevelUp,
		MatchEventTypeAscendedEvent,
		MatchEventTypeCapturePoint,
		MatchEventTypePoroKingSummon,
	}
)

// MatchEvent is an event in a match at a certain timestamp
type MatchEvent struct {
	EventType               string          `json:"eventType"`
	TowerType               string          `json:"towerType"`
	TeamID                  int             `json:"teamId"`
	AscendedType            string          `json:"ascendedType"`
	KillerID                int             `json:"killerId"`
	LevelUpType             string          `json:"levelUpType"`
	PointCaptured           string          `json:"pointCaptured"`
	AssistingParticipantIDs []int           `json:"assistingParticipantIds"`
	WardType                string          `json:"wardType"`
	MonsterType             string          `json:"monsterType"`
	Type                    *MatchEventType `json:"type"`
	SkillSlot               int             `json:"skillSlot"`
	VictimID                int             `json:"victimId"`
	Timestamp               int             `json:"timestamp"`
	AfterID                 int             `json:"afterId"`
	MonsterSubType          string          `json:"monsterSubType"`
	LaneType                string          `json:"laneType"`
	ItemID                  int             `json:"itemId"`
	ParticipantID           int             `json:"participantId"`
	BuildingType            string          `json:"buildingType"`
	CreatorID               int             `json:"creatorId"`
	Position                *MatchPosition  `json:"position"`
	BeforeID                int             `json:"beforeId"`
}

// GetItem returns the item for this event
func (e *MatchEvent) GetItem(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(e.ItemID))
}

// MatchPosition is a position on the map in a game
type MatchPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// GameInfo contains information about an ongoing game
type GameInfo struct {
	GameID            int                       `json:"gameId"`
	GameStartTime     int                       `json:"gameStartTime"`
	PlatformID        string                    `json:"platformId"`
	GameMode          string                    `json:"gameMode"`
	MapID             int                       `json:"mapId"`
	GameType          string                    `json:"gameType"`
	BannedChampions   []*BannedChampion         `json:"bannedChampions"`
	Observers         *Observer                 `json:"observers"`
	Participants      []*CurrentGameParticipant `json:"participants"`
	GameLength        int                       `json:"gameLength"`
	GameQueueConfigID int                       `json:"gameQueueConfigId"`
}

// GetMatch returns information about the finished match
func (i *GameInfo) GetMatch(client *Client) (*Match, error) {
	return client.Match.Get(fmt.Sprintf("%v_%v", strings.ToUpper(string(client.Match.c.Region)), i.GameID))
}

// BannedChampion represents a champion ban during pack/ban phase
type BannedChampion struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
	TeamID     int `json:"teamId"`
}

// GetChampion returns the banned champion
func (c *BannedChampion) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(c.ChampionID))
}

// Observer is an observer of an ongoing game
type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

// CurrentGameParticipant represents a player in an ongoing game
type CurrentGameParticipant struct {
	ProfileIconID            int                        `json:"profileIconId"`
	ChampionID               int                        `json:"championId"`
	SummonerName             string                     `json:"summonerName"`
	GameCustomizationObjects []*GameCustomizationObject `json:"gameCustomizationObjects"`
	Bot                      bool                       `json:"bot"`
	Perks                    *Perks                     `json:"perks"`
	Spell2ID                 int                        `json:"spell2Id"`
	Spell1ID                 int                        `json:"spell1Id"`
	TeamID                   int                        `json:"teamId"`
	SummonerID               string                     `json:"summonerId"`
}

// GetChampion returns the champion played by this participant
func (p *CurrentGameParticipant) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(p.ChampionID))
}

// GetSpell1 returns the first summoner spell of this participant
func (p *CurrentGameParticipant) GetSpell1(client *datadragon.Client) (datadragon.SummonerSpell, error) {
	return client.GetSummonerSpell(strconv.Itoa(p.Spell1ID))
}

// GetSpell2 returns the second summoner spell of this participant
func (p *CurrentGameParticipant) GetSpell2(client *datadragon.Client) (datadragon.SummonerSpell, error) {
	return client.GetSummonerSpell(strconv.Itoa(p.Spell2ID))
}

// GameCustomizationObject contains information specific to an ongoing game
type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}

// Perks represents the runes for a player in an ongoing game
type Perks struct {
	PerkStyle    int   `json:"perkStyle"`
	PerksIDs     []int `json:"perkIds"`
	PerkSubStyle int   `json:"perkSubStyle"`
}

// FeaturedGames represents a list of featured games
type FeaturedGames struct {
	ClientRefreshInterval int         `json:"clientRefreshInterval"`
	GameList              []*GameInfo `json:"gameList"`
}

// Status contains information about all services in a certain region
type Status struct {
	Name      string     `json:"name"`
	RegionTag string     `json:"region_tag"`
	Hostname  string     `json:"hostname"`
	Services  []*Service `json:"services"`
	Slug      string     `json:"slug"`
	Locales   []string   `json:"locales"`
}

// Service is a service provided by Riot with its status
type Service struct {
	Status    string      `json:"status"`
	Incidents []*Incident `json:"incidents"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
}

// Incident contains information about an incident
type Incident struct {
	Active    bool             `json:"active"`
	CreatedAt string           `json:"created_at"`
	ID        int              `json:"id"`
	Updates   []*StatusMessage `json:"updates"`
}

// StatusMessage contains information about a status message
type StatusMessage struct {
	Severity     string               `json:"severity"`
	Author       string               `json:"author"`
	CreatedAt    string               `json:"created_at"`
	Translations []*StatusTranslation `json:"translations"`
	UpdatedAt    string               `json:"updated_at"`
	Content      string               `json:"content"`
	ID           string               `json:"id"`
}

// StatusTranslation contains the status message content in a certain language
type StatusTranslation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

// Summoner represents a summoner with several related IDs
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	PUUID         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	RevisionDate  int    `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}

// LobbyEventList is a wrapper for a list of lobby events in a tournament
type LobbyEventList struct {
	EventList []*LobbyEvent `json:"eventList"`
}

// LobbyEvent represents an event that happened in a tournament lobby
type LobbyEvent struct {
	EventType  string `json:"eventType"`
	SummonerID string `json:"summonerId"`
	Timestamp  string `json:"timestamp"`
}

// Tournament contains the settings of a previously created tournament
type Tournament struct {
	Map          string   `json:"map"`
	Code         string   `json:"code"`
	Spectators   string   `json:"spectators"`
	Region       string   `json:"region"`
	ProviderID   int      `json:"providerId"`
	TeamSize     int      `json:"teamSize"`
	Participants []string `json:"participants"`
	PickType     string   `json:"pickType"`
	TournamentID int      `json:"tournamentId"`
	LobbyName    string   `json:"lobbyName"`
	Password     string   `json:"password"`
	ID           int      `json:"id"`
	MetaData     string   `json:"metaData"`
}

// TournamentCodeParameters parameters needed to create tournament codes
type TournamentCodeParameters struct {
	// The spectator type of the game. (Legal values: NONE, LOBBYONLY, ALL)
	SpectatorType string `json:"spectatorType"`
	// The team size of the game. Valid values are 1-5.
	TeamSize int `json:"teamSize"`
	// The pick type of the game. (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT)
	PickType string `json:"pickType"`
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby.
	// NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and
	// teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedSummonerIDs []string `json:"allowedSummonerIds,omitempty"`
	// The map type of the game. (Legal values: SUMMONERS_RIFT, TWISTED_TREELINE, HOWLING_ABYSS)
	MapType string `json:"mapType"`
	// Optional string that may contain any data in any format, if specified at all. Used to denote any custom
	// information about the game.
	Metadata string `json:"metadata"`
}

// TournamentUpdateParameters parameters needed to update an existing tournament
type TournamentUpdateParameters struct {
	// The spectator type (Legal values: NONE, LOBBYONLY, ALL)
	SpectatorType string `json:"spectatorType"`
	// The pick type (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT)
	PickType string `json:"pickType"`
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby.
	// NOTE: Participants are not enforced at the team level, but rather the aggregate of teamOne and teamTwo.
	AllowedSummonerIDs []string `json:"allowedSummonerIds"`
	// The map type (Legal values: SUMMONERS_RIFT, TWISTED_TREELINE, HOWLING_ABYSS)
	MapType string `json:"mapType"`
}

// TournamentRegistrationParameters parameters required for creating a tournament
type TournamentRegistrationParameters struct {
	// The provider ID to specify the regional registered provider data to associate this tournament.
	ProviderID int `json:"providerId"`
	// The optional name of the tournament.
	Name string `json:"name"`
}

// ProviderRegistrationParameters parameters required for registering a provider with tournaments for a region
type ProviderRegistrationParameters struct {
	// The provider's callback URL to which tournament game results in this region should be posted. The URL must be
	// well-formed, use the http or https protocol, and use the default port for the protocol (http URLs must use port
	// 80, https URLs must use port 443).
	URL string `json:"url"`
	// The region in which the provider will be running tournaments.
	// (Legal values: BR, EUNE, EUW, JP, LAN, LAS, NA, OCE, PBE, RU, TR)
	Region string `json:"region"`
}

// ChallengeConfigInfo represents basic challenge configuration information
type ChallengeConfigInfo struct {
	ID             int64                        `json:"id"`
	LocalizedNames map[string]map[string]string `json:"localizedNames"`
	State          string                       `json:"state"`
	Tracking       string                       `json:"tracking"`
	StartTimeStamp int64                        `json:"startTimeStamp"`
	EndTimeStamp   int64                        `json:"endTimeStamp"`
	Leaderboard    bool                         `json:"leaderboard"`
	Thresholds     map[string]float64           `json:"thresholds"`
}

// ChallengePoints contains the settings of a previously created tournament
type ChallengePoints struct {
	Level      string  `json:"level"`
	Current    float32 `json:"current"`
	Max        int32   `json:"max"`
	Percentile float32 `json:"percentile"`
}

// ChallengeInfo represents each challenge info for a player
type ChallengeInfo struct {
	ChallengeID  int32   `json:"challengeid"`
	Percentile   float32 `json:"percentile"`
	Level        string  `json:"level"`
	Value        float32 `json:"value"`
	AchievedTime int64   `json:"achievedtime"`
}

// PlayerClientPreferences holds player preferences
type PlayerClientPreferences struct {
	BannerAccent string  `json:"banneraccent"`
	Title        string  `json:"title"`
	ChallengeID  []int32 `json:"challengeids"`
}

// PlayerInfo contains player information with list of all progressed challenges
type PlayerInfo struct {
	TotalPoints    *ChallengePoints           `json:"totalpoints"`
	CategoryPoints map[string]ChallengePoints `json:"categorypoints"`
	Challenges     []*ChallengeInfo           `json:"challenges"`
	Preferences    *PlayerClientPreferences   `json:"preferences"`
}

// ApexPlayerInfo holds information of top players for each level
type ApexPlayerInfo struct {
	PuuID    string  `json:"puuid"`
	Value    float64 `json:"value"`
	Position int32   `json:"position"`
}
