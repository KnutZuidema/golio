package tft

const (
	endpointBase          = "/tft"
	endpointSpectatorBase = "/lol/spectator" + endpointBase

	endpointSpectatorActiveGamedByPUUID = endpointSpectatorBase + "/v5/active-games/by-puuid/%s"
	endpointSpectatorFeaturedGames      = endpointSpectatorBase + "/v5/featured-games"

	endpointLeagueBase                = endpointBase + "/league/v1"
	endpointLeagueChallenger          = endpointLeagueBase + "/challenger?queue=%s"
	endpointLeagueEntriesBySummoner   = endpointLeagueBase + "/entries/by-summoner/%s"
	endpointLeagueEntries             = endpointLeagueBase + "/entries/%s/%s"
	endpointLeagueGrandMaster         = endpointLeagueBase + "/grandmaster?queue=%s"
	endpointLeagueLeagues             = endpointLeagueBase + "/leagues/%s"
	endpointLeagueMaster              = endpointLeagueBase + "/master?queue=%s"
	endpointLeagueRatedLattersByQueue = endpointLeagueBase + "/rated-ladders/%s/top"

	endpointMatchBase      = endpointBase + "/match/v1/matches"
	endpointMatchesByPUUID = endpointMatchBase + "/by-puuid/%s/ids"
	endpointMatchByMatchID = endpointMatchBase + "/%s"

	endpointStatusBase         = endpointBase + "/status/v1"
	endpointStatusPlatformData = endpointStatusBase + "/platform-data"

	endpointSummonerBase         = endpointBase + "/summoner/v1/summoners"
	endpointSummonerByAccount    = endpointSummonerBase + "/by-account/%s"
	endpointSummonerByPUUID      = endpointSummonerBase + "/by-puuid/%s"
	endpointSummonerByMe         = endpointSummonerBase + "/me"
	endpointSummonerBySummonerID = endpointSummonerBase + "/%s"
)

type queue string

const (
	QueueRankedTFT         queue = "RANKED_TFT"
	QueueRankedTFTDoubleUp queue = "RANKED_TFT_DOUBLE_UP"
	QueueRankedTFTTurbo    queue = "RANKED_TFT_TURBO"
)

type tier string

// All possible Tiers
const (
	TierIron        tier = "IRON"
	TierBronze      tier = "BRONZE"
	TierSilver      tier = "SILVER"
	TierGold        tier = "GOLD"
	TierPlatinum    tier = "PLATINUM"
	TierEmerald     tier = "EMERALD"
	TierDiamond     tier = "DIAMOND"
	TierMaster      tier = "MASTER"
	TierGrandMaster tier = "GRANDMASTER"
	TierChallenger  tier = "CHALLENGER"
)

type division string

// All possible divisions
const (
	DivisionOne   division = "I"
	DivisionTwo   division = "II"
	DivisionThree division = "III"
	DivisionFour  division = "IV"
)
