package tft

const (
	endpointBase                         = "/tft"
	endpointMasteryBase                  = endpointBase + "/champion-mastery/v4"
	endpointGetChampionMasteries         = endpointMasteryBase + "/champion-masteries/by-summoner/%s"
	endpointGetChampionMastery           = endpointMasteryBase + "/champion-masteries/by-summoner/%s/by-champion/%s"
	endpointGetChampionMasteryTotalScore = endpointMasteryBase + "/scores/by-summoner/%s"
	endpointPlatformBase                 = endpointBase + "/platform/v3"
	endpointGetFreeChampionRotation      = endpointPlatformBase + "/champion-rotations"
	endpointLeagueBase                   = endpointBase + "/league/v1"
	endpointGetChallengerLeague          = endpointLeagueBase + "/challengerleagues/by-queue/%s"
	endpointGetGrandmasterLeague         = endpointLeagueBase + "/grandmasterleagues/by-queue/%s"
	endpointGetMasterLeague              = endpointLeagueBase + "/masterleagues/by-queue/%s"
	endpointGetLeaguesBySummoner         = endpointLeagueBase + "/entries/by-summoner/%s"
	endpointGetLeagues                   = endpointLeagueBase + "/entries/%s/%s/%s"
	endpointGetLeague                    = endpointLeagueBase + "/leagues/%s"
	endpointStatusBase                   = endpointBase + "/status/v3"
	endpointGetStatus                    = endpointStatusBase + "/shard-data"
	endpointMatchBase                    = endpointBase + "/match/v4"
	endpointGetMatch                     = endpointMatchBase + "/matches/%d"
	endpointGetMatchesByAccount          = endpointMatchBase + "/matchlists/by-account/%s?beginIndex=%d&endIndex=%d"
	endpointGetMatchTimeline             = endpointMatchBase + "/timelines/by-match/%d"
	endpointGetMatchIDsByTournamentCode  = endpointMatchBase + "/matches/by-tournament-code/%s/ids"
	endpointGetMatchForTournament        = endpointMatchBase + "/matches/%d/by-tournament-code/%s"
	endpointSummonerBase                 = endpointBase + "/summoner/v1"
	endpointGetSummonerBySummonerID      = endpointSummonerBase + "/summoners/%s"
	endpointGetSummonerBy                = endpointSummonerBase + "/summoners/by-%s/%s"
	endpointSpectatorBase                = endpointBase + "/spectator/v4"
	endpointGetCurrentGame               = endpointSpectatorBase + "/active-games/by-summoner/%s"
	endpointGetFeaturedGames             = endpointSpectatorBase + "/featured-games"
	endpointTournamentStubBase           = endpointBase + "/tournament-stub/v4"
	endpointCreateStubTournamentCodes    = endpointTournamentStubBase + "/codes?count=%d&tournamentId=%d"
	endpointGetStubLobbyEvents           = endpointTournamentStubBase + "/lobby-events/by-code/%s"
	endpointCreateStubTournamentProvider = endpointTournamentStubBase + "/providers"
	endpointCreateStubTournament         = endpointTournamentStubBase + "/tournaments"
	endpointTournamentBase               = endpointBase + "/tournament/v4"
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
	QueueRankedTfT queue = "RANKED_TFT"
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
		QueueRankedTfT,
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
