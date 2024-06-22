package tft

const (
	endpointBase          = "/tft"
	endpointSpectatorBase = "/lol/spectator" + endpointBase

	endpointSpectatorActiveGamedByPUUID = endpointSpectatorBase + "/v5/active-games/by-puuid/%s"
	endpointSpectatorFeaturedGames      = endpointSpectatorBase + "/v5/featured-games"

	endpointLeagueBase                = "/league/v1"
	endpointLeagueChallenger          = endpointLeagueBase + "/challenger"
	endpointLeagueEntriesBySummoner   = endpointLeagueBase + "/entries/by-summoner/%s"
	endpointLeagueEntries             = endpointLeagueBase + "/entries/%s/%s"
	endpointLeagueGrandMaster         = endpointLeagueBase + "/grandmaster"
	endpointLeagueLeagues             = endpointLeagueBase + "/leagues/%s"
	endpointLeagueMaster              = endpointLeagueBase + "/master"
	endpointLeagueRatedLattersByQueue = endpointLeagueBase + "/rated-ladders/%s/top"
)
