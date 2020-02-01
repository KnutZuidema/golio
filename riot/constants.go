package riot

const (
	apiURLFormat                         = "%s://%s.%s%s"
	baseURL                              = "api.riotgames.com"
	scheme                               = "https"
	apiTokenHeaderKey                    = "X-Riot-Token"
	endpointMasteryBase                  = "/riot/champion-mastery/v4"
	endpointGetChampionMasteries         = endpointMasteryBase + "/champion-masteries/by-summoner/%s"
	endpointGetChampionMastery           = endpointMasteryBase + "/champion-masteries/by-summoner/%s/by-champion/%s"
	endpointGetChampionMasteryTotalScore = endpointMasteryBase + "/scores/by-summoner/%s"
	endpointPlatformBase                 = "/riot/platform/v3"
	endpointGetFreeChampionRotation      = endpointPlatformBase + "/champion-rotations"
	endpointLeagueBase                   = "/riot/league/v4"
	endpointGetChallengerLeague          = endpointLeagueBase + "/challengerleagues/by-queue/%s"
	endpointGetGrandmasterLeague         = endpointLeagueBase + "/grandmasterleagues/by-queue/%s"
	endpointGetMasterLeague              = endpointLeagueBase + "/masterleagues/by-queue/%s"
	endpointGetLeaguesBySummoner         = endpointLeagueBase + "/entries/by-summoner/%s"
	endpointGetLeagues                   = endpointLeagueBase + "/entries/%s/%s/%s"
	endpointGetLeague                    = endpointLeagueBase + "/leagues/%s"
	endpointStatusBase                   = "/riot/status/v3"
	endpointGetStatus                    = endpointStatusBase + "/shard-data"
	endpointMatchBase                    = "/riot/match/v4"
	endpointGetMatch                     = endpointMatchBase + "/matches/%d"
	endpointGetMatchesByAccount          = endpointMatchBase + "/matchlists/by-account/%s?beginIndex=%d&endIndex=%d"
	endpointGetMatchTimeline             = endpointMatchBase + "/timelines/by-match/%d"
	endpointGetMatchIDsByTournamentCode  = endpointMatchBase + "/matches/by-tournament-code/%s/ids"
	endpointGetMatchForTournament        = endpointMatchBase + "/matches/%d/by-tournament-code/%s"
	endpointSummonerBase                 = "/riot/summoner/v4"
	endpointGetSummonerBySummonerID      = endpointSummonerBase + "/summoners/%s"
	endpointGetSummonerBy                = endpointSummonerBase + "/summoners/by-%s/%s"
	endpointSpectatorBase                = "/riot/spectator/v4"
	endpointGetCurrentGame               = endpointSpectatorBase + "/active-games/by-summoner/%s"
	endpointGetFeaturedGames             = endpointSpectatorBase + "/featured-games"
	endpointTournamentStubBase           = "/riot/tournament-stub/v4"
	endpointCreateStubTournamentCodes    = endpointTournamentStubBase + "/codes?count=%d&tournamentId=%d"
	endpointGetStubLobbyEvents           = endpointTournamentStubBase + "/lobby-events/by-code/%s"
	endpointCreateStubTournamentProvider = endpointTournamentStubBase + "/providers"
	endpointCreateStubTournament         = endpointTournamentStubBase + "/tournaments"
	endpointTournamentBase               = "/riot/tournament/v4"
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
)
