package riot

import (
	"sort"
	"strconv"

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
		sort.Slice(l.sortedEntries, func(i, j int) bool {
			return l.sortedEntries[i].LeaguePoints > l.sortedEntries[j].LeaguePoints
		})
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
	// Please refer to the Game Constants documentation.
	SeasonID int `json:"seasonId"`
	// Please refer to the Game Constants documentation.
	QueueID int `json:"queueId"`
	GameID  int `json:"gameId"`
	// Participant identity information.
	ParticipantIdentities []*ParticipantIdentity `json:"participantIdentities"`
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
	Teams []*TeamStats `json:"teams"`
	// Participant information.
	Participants []*Participant `json:"participants"`
	// Match duration in seconds.
	GameDuration int `json:"gameDuration"`
	// Designates the timestamp when champion select ended and the loading screen appeared, NOT when the game timer was
	// at 0:00.
	GameCreation int `json:"gameCreation"`
}

// GetSeason returns the season this match was played in
func (m *Match) GetSeason(client *static.Client) (static.Season, error) {
	return client.GetSeason(m.SeasonID)
}

// GetQueue returns the queue this match was played in
func (m *Match) GetQueue(client *static.Client) (static.Queue, error) {
	return client.GetQueue(m.QueueID)
}

// GetMap returns the map this match was played on
func (m *Match) GetMap(client *static.Client) (static.Map, error) {
	return client.GetMap(m.MapID)
}

// GetGameType returns the gameType this match was played in
func (m *Match) GetGameType(client *static.Client) (static.GameType, error) {
	return client.GetGameType(m.GameType)
}

// GetGameMode returns the gameMode this match was played in
func (m *Match) GetGameMode(client *static.Client) (static.GameMode, error) {
	return client.GetGameMode(m.GameMode)
}

// ParticipantIdentity contains a reference to a player for a participant in a game
type ParticipantIdentity struct {
	// Player information.
	Player        *Player `json:"player"`
	ParticipantID int     `json:"participantId"`
}

// Player represents a player
type Player struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
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

// GetSummoner returns the summoner info for this player
func (p *Player) GetSummoner(client *Client) (*Summoner, error) {
	return client.Summoner.GetByID(p.SummonerID)
}

// GetProfileIcon returns the profile icon data for this player
func (p *Player) GetProfileIcon(client *datadragon.Client) (datadragon.ProfileIcon, error) {
	return client.GetProfileIcon(p.ProfileIcon)
}

// TeamStats represents the stats of a team for a single game
type TeamStats struct {
	// Flag indicating whether or not the team scored the first Dragon kill.
	FirstDragon bool `json:"firstDragon"`
	// Flag indicating whether or not the team destroyed the first inhibitor.
	FirstInhibitor bool `json:"firstInhibitor"`
	// If match queueID has a draft, contains banned champion data, otherwise empty.
	Bans []*TeamBan `json:"bans"`
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
	// string indicating whether or not the team won. There are only two values visibile in public match history.
	// (Legal values: Fail, Win)
	Win string `json:"win"`
	// Number of times the team killed Dragon.
	DragonKills int `json:"dragonKills"`
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

// Participant represents a participant in a game
type Participant struct {
	// Participant statistics.
	Stats         *ParticipantStats `json:"stats"`
	ParticipantID int               `json:"participantId"`
	// List of legacy Rune information. Not included for matches played with Runes Reforged.
	Runes []*Rune `json:"runes"`
	// Participant timeline data.
	Timeline *ParticipantTimeline `json:"timeline"`
	// 100 for blue side. 200 for red side.
	TeamID int `json:"teamId"`
	// Second Summoner Spell id.
	Spell2ID int `json:"spell2Id"`
	// List of legacy Mastery information. Not included for matches played with Runes Reforged.
	Masteries []*LegacyMastery `json:"masteries"`
	// Highest ranked tier achieved for the previous season in a specific subset of queueIds, if any, otherwise null.
	// Used to display border in game loading screen. Please refer to the Ranked Info documentation.
	// (Legal values: CHALLENGER, MASTER, DIAMOND, PLATINUM, GOLD, SILVER, BRONZE, UNRANKED)
	HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
	// First Summoner Spell id.
	Spell1ID   int `json:"spell1Id"`
	ChampionID int `json:"championId"`
}

// GetChampion returns the champion played by this participant
func (p *Participant) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(p.ChampionID))
}

// GetSpell1 returns the first summoner spell of this participant
func (p *Participant) GetSpell1(client *datadragon.Client) (datadragon.SummonerSpell, error) {
	return client.GetSummonerSpell(strconv.Itoa(p.Spell1ID))
}

// GetSpell2 returns the second summoner spell of this participant
func (p *Participant) GetSpell2(client *datadragon.Client) (datadragon.SummonerSpell, error) {
	return client.GetSummonerSpell(strconv.Itoa(p.Spell2ID))
}

// ParticipantStats contains stats of a participant in a game
type ParticipantStats struct {
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	VisionScore                     int  `json:"visionScore"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	IntestTimeSpentLiving           int  `json:"intestTimeSpentLiving"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	NodeCapture                     int  `json:"nodeCapture"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	WardsKilled                     int  `json:"wardsKilled"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	TeamObjective                   int  `json:"teamObjective"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
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
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	TurretKills                     int  `json:"turretKills"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	GoldEarned                      int  `json:"goldEarned"`
	KillingSprees                   int  `json:"killingSprees"`
	UnrealKills                     int  `json:"unrealKills"`
	AltarsCaptured                  int  `json:"altarsCaptured"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	ChampLevel                      int  `json:"champLevel"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	AltarsNeutralized               int  `json:"altarsNeutralized"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`

	// Primary rune path
	PerkPrimaryStyle int `json:"perkPrimaryStyle"`
	// Secondary rune path
	PerkSubStyle int `json:"perkSubStyle"`
	// Primary path keystone rune.
	Perk0     int `json:"perk0"`
	Perk0Var1 int `json:"perk0Var1"`
	Perk0Var2 int `json:"perk0Var2"`
	Perk0Var3 int `json:"perk0Var3"`
	// Primary path rune.
	Perk1     int `json:"perk1"`
	Perk1Var1 int `json:"perk1Var1"`
	Perk1Var2 int `json:"perk1Var2"`
	Perk1Var3 int `json:"perk1Var3"`
	// Primary path rune.
	Perk2     int `json:"perk2"`
	Perk2Var1 int `json:"perk2Var1"`
	Perk2Var2 int `json:"perk2Var2"`
	Perk2Var3 int `json:"perk2Var3"`
	// Primary path rune.
	Perk3     int `json:"perk3"`
	Perk3Var1 int `json:"perk3Var1"`
	Perk3Var2 int `json:"perk3Var2"`
	Perk3Var3 int `json:"perk3Var3"`
	// Secondary path rune.
	Perk4     int `json:"perk4"`
	Perk4Var1 int `json:"perk4Var1"`
	Perk4Var2 int `json:"perk4Var2"`
	Perk4Var3 int `json:"perk4Var3"`
	// Secondary path rune.
	Perk5     int `json:"perk5"`
	Perk5Var1 int `json:"perk5Var1"`
	Perk5Var2 int `json:"perk5Var2"`
	Perk5Var3 int `json:"perk5Var3"`

	PlayerScore0 int `json:"playerScore0"`
	PlayerScore1 int `json:"playerScore1"`
	PlayerScore2 int `json:"playerScore2"`
	PlayerScore3 int `json:"playerScore3"`
	PlayerScore4 int `json:"playerScore4"`
	PlayerScore5 int `json:"playerScore5"`
	PlayerScore6 int `json:"playerScore6"`
	PlayerScore7 int `json:"playerScore7"`
	PlayerScore8 int `json:"playerScore8"`
	PlayerScore9 int `json:"playerScore9"`

	Item0 int `json:"item0"`
	Item1 int `json:"item1"`
	Item2 int `json:"item2"`
	Item3 int `json:"item3"`
	Item4 int `json:"item4"`
	Item5 int `json:"item5"`
	Item6 int `json:"item6"`

	Deaths      int `json:"deaths"`
	Assists     int `json:"assists"`
	Kills       int `json:"kills"`
	DoubleKills int `json:"doubleKills"`
	TripleKills int `json:"tripleKills"`
	QuadraKills int `json:"quadraKills"`
	PentaKills  int `json:"pentaKills"`
}

// GetItem0 returns the item in slot 0
func (s *ParticipantStats) GetItem0(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item0))
}

// GetItem1 returns the item in slot 1
func (s *ParticipantStats) GetItem1(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item1))
}

// GetItem2 returns the item in slot 2
func (s *ParticipantStats) GetItem2(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item2))
}

// GetItem3 returns the item in slot 3
func (s *ParticipantStats) GetItem3(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item3))
}

// GetItem4 returns the item in slot 4
func (s *ParticipantStats) GetItem4(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item4))
}

// GetItem5 returns the item in slot 5
func (s *ParticipantStats) GetItem5(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item5))
}

// GetItem6 returns the item in slot 6
func (s *ParticipantStats) GetItem6(client *datadragon.Client) (datadragon.Item, error) {
	return client.GetItem(strconv.Itoa(s.Item6))
}

// Rune represents an old rune
type Rune struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

// ParticipantTimeline contains timeline values for a participant in a game
type ParticipantTimeline struct {
	// Participant's calculated lane. MID and BOT are legacy values.
	// (Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM)
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

// LegacyMastery represents the old masteries
type LegacyMastery struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

// Matchlist contains information about all games played by a single summoner
type Matchlist struct {
	Matches    []*MatchReference `json:"matches"`
	TotalGames int               `json:"totalGames"`
	StartIndex int               `json:"startIndex"`
	EndIndex   int               `json:"endIndex"`
}

// MatchReference contains information about a game by a single summoner
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

// GetChampion returns the champion played in this match
func (r *MatchReference) GetChampion(client *datadragon.Client) (datadragon.ChampionDataExtended, error) {
	return client.GetChampionByID(strconv.Itoa(r.Champion))
}

// GetSeason returns the season this match as played in
func (r *MatchReference) GetSeason(client *static.Client) (static.Season, error) {
	return client.GetSeason(r.Season)
}

// GetQueue returns the queue this match was played in
func (r *MatchReference) GetQueue(client *static.Client) (static.Queue, error) {
	return client.GetQueue(r.Queue)
}

// GetGame returns more information about this match
func (r *MatchReference) GetGame(client *Client) (*Match, error) {
	return client.Match.Get(r.GameID)
}

// MatchTimeline contains timeline frames for a match
type MatchTimeline struct {
	Frames   []*MatchFrame `json:"frames"`
	Interval int           `json:"interval"`
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
	return client.Match.Get(i.GameID)
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
	PerksIDs     []int `json:"perksIDs"`
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
	AllowedSummonerIDs []string `json:"allowedSummonerIds"`
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
