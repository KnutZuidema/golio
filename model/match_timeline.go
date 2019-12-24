package model

// MatchTimeline contains timeline frames for a match
type MatchTimeline struct {
	Frames   []MatchFrame `json:"frames"`
	Interval int          `json:"interval"`
}

// MatchFrame is a single frame in the timeline of a game
type MatchFrame struct {
	Timestamp         int                         `json:"timestamp"`
	ParticipantFrames map[string]ParticipantFrame `json:"participantFrames"`
	Events            []MatchEvent                `json:"events"`
}

// ParticipantFrame contains information about a participant in a game at a single timestamp
type ParticipantFrame struct {
	TotalGold           int           `json:"totalGold"`
	TeamScore           int           `json:"teamScore"`
	ParticipantID       int           `json:"participantId"`
	Level               int           `json:"level"`
	CurrentGold         int           `json:"currentGold"`
	MinionsKilled       int           `json:"minionsKilled"`
	DominionScore       int           `json:"dominionScore"`
	Position            MatchPosition `json:"position"`
	XP                  int           `json:"xp"`
	JungleMinionsKilled int           `json:"jungleMinionsKilled"`
}

// MatchEvent is an event in a match at a certain timestamp
type MatchEvent struct {
	EventType               string         `json:"eventType"`
	TowerType               string         `json:"towerType"`
	TeamID                  int            `json:"teamId"`
	AscendedType            string         `json:"ascendedType"`
	KillerID                int            `json:"killerId"`
	LevelUpType             string         `json:"levelUpType"`
	PointCaptured           string         `json:"pointCaptured"`
	AssistingParticipantIDs []int          `json:"assistingParticipantIds"`
	WardType                string         `json:"wardType"`
	MonsterType             string         `json:"monsterType"`
	Type                    MatchEventType `json:"type"`
	SkillSlot               int            `json:"skillSlot"`
	VictimID                int            `json:"victimId"`
	Timestamp               int            `json:"timestamp"`
	AfterID                 int            `json:"afterId"`
	MonsterSubType          string         `json:"monsterSubType"`
	LaneType                string         `json:"laneType"`
	ItemID                  int            `json:"itemId"`
	ParticipantID           int            `json:"participantId"`
	BuildingType            string         `json:"buildingType"`
	CreatorID               int            `json:"creatorId"`
	Position                MatchPosition  `json:"position"`
	BeforeID                int            `json:"beforeId"`
}

// MatchPosition is a position on the map in a game
type MatchPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
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
