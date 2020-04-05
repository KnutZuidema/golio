package riot

const (
	apiURLFormat                         = "%s://%s.%s%s"
	baseURL                              = "api.riotgames.com"
	scheme                               = "https"
	apiTokenHeaderKey                    = "X-Riot-Token"
	endpointMasteryBase                  = "/lol/champion-mastery/v4"
	endpointGetChampionMasteries         = endpointMasteryBase + "/champion-masteries/by-summoner/%s"
	endpointGetChampionMastery           = endpointMasteryBase + "/champion-masteries/by-summoner/%s/by-champion/%s"
	endpointGetChampionMasteryTotalScore = endpointMasteryBase + "/scores/by-summoner/%s"
	endpointPlatformBase                 = "/lol/platform/v3"
	endpointGetFreeChampionRotation      = endpointPlatformBase + "/champion-rotations"
	endpointLeagueBase                   = "/lol/league/v4"
	endpointGetChallengerLeague          = endpointLeagueBase + "/challengerleagues/by-queue/%s"
	endpointGetGrandmasterLeague         = endpointLeagueBase + "/grandmasterleagues/by-queue/%s"
	endpointGetMasterLeague              = endpointLeagueBase + "/masterleagues/by-queue/%s"
	endpointGetLeaguesBySummoner         = endpointLeagueBase + "/entries/by-summoner/%s"
	endpointGetLeagues                   = endpointLeagueBase + "/entries/%s/%s/%s"
	endpointGetLeague                    = endpointLeagueBase + "/leagues/%s"
	endpointStatusBase                   = "/lol/status/v3"
	endpointGetStatus                    = endpointStatusBase + "/shard-data"
	endpointMatchBase                    = "/lol/match/v4"
	endpointGetMatch                     = endpointMatchBase + "/matches/%d"
	endpointGetMatchesByAccount          = endpointMatchBase + "/matchlists/by-account/%s?beginIndex=%d&endIndex=%d"
	endpointGetMatchTimeline             = endpointMatchBase + "/timelines/by-match/%d"
	endpointGetMatchIDsByTournamentCode  = endpointMatchBase + "/matches/by-tournament-code/%s/ids"
	endpointGetMatchForTournament        = endpointMatchBase + "/matches/%d/by-tournament-code/%s"
	endpointSummonerBase                 = "/lol/summoner/v4"
	endpointGetSummonerBySummonerID      = endpointSummonerBase + "/summoners/%s"
	endpointGetSummonerBy                = endpointSummonerBase + "/summoners/by-%s/%s"
	endpointSpectatorBase                = "/lol/spectator/v4"
	endpointGetCurrentGame               = endpointSpectatorBase + "/active-games/by-summoner/%s"
	endpointGetFeaturedGames             = endpointSpectatorBase + "/featured-games"
	endpointTournamentStubBase           = "/lol/tournament-stub/v4"
	endpointCreateStubTournamentCodes    = endpointTournamentStubBase + "/codes?count=%d&tournamentId=%d"
	endpointGetStubLobbyEvents           = endpointTournamentStubBase + "/lobby-events/by-code/%s"
	endpointCreateStubTournamentProvider = endpointTournamentStubBase + "/providers"
	endpointCreateStubTournament         = endpointTournamentStubBase + "/tournaments"
	endpointTournamentBase               = "/lol/tournament/v4"
	endpointCreateTournamentCodes        = endpointTournamentBase + "/codes?count=%d&tournamentId=%d"
	endpointGetLobbyEvents               = endpointTournamentBase + "/lobby-events/by-code/%s"
	endpointCreateTournamentProvider     = endpointTournamentBase + "/providers"
	endpointCreateTournament             = endpointTournamentBase + "/tournaments"
	endpointGetTournament                = endpointTournamentBase + "/codes/%s"
	endpointUpdateTournament             = endpointTournamentBase + "/codes/%s"
	endpointGetThirdPartyCode            = endpointPlatformBase + "/third-party-code/by-summoner/%s"
)

type identification string

const (
	identificationName       identification = "name"
	identificationAccountID                 = "account"
	identificationPUUID                     = "puuid"
	identificationSummonerID                = "summonerID"
)

type queue string

// All possible queues
const (
	QueueRankedSolo            queue = "RANKED_SOLO_5x5"
	QueueRankedFlex                  = "RANKED_FLEX_SR"
	QueueRankedTwistedTreeline       = "RANKED_FLEX_TT"
)

type tier string

// All possible Tiers
const (
	TierIron     tier = "IRON"
	TierBronze        = "BRONZE"
	TierSilver        = "SILVER"
	TierGold          = "GOLD"
	TierPlatinum      = "PLATINUM"
	TierDiamond       = "DIAMOND"
)

type division string

// All possible divisions
const (
	DivisionOne   division = "I"
	DivisionTwo            = "II"
	DivisionThree          = "III"
	DivisionFour           = "IV"
)

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
	// Queues is a list of all available queue types
	Queues = []queue{
		QueueRankedSolo,
		QueueRankedFlex,
		QueueRankedTwistedTreeline,
	}

	// Tiers is a list of all available tiers
	Tiers = []tier{
		TierIron,
		TierBronze,
		TierSilver,
		TierGold,
		TierPlatinum,
		TierDiamond,
	}

	// Divisions is a list of all available divisions
	Divisions = []division{
		DivisionOne,
		DivisionTwo,
		DivisionThree,
		DivisionFour,
	}

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
