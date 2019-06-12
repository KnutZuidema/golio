package api

// Region represents a server region
type Region string
type identification string
type queue string
type tier string
type division string
type dataDragonURL string
type languageCode string

// All possible language codes
const (
	LanguageCodeCzechRepublic            languageCode = "cs_CZ"
	LanguageCodeGreece                                = "el_GR"
	LanguageCodePoland                                = "pl_PL"
	LanguageCodeRomania                               = "ro_RO"
	LanguageCodeHungary                               = "hu_HU"
	LanguageCodeUnitedKingdom                         = "en_GB"
	LanguageCodeGermany                               = "de_DE"
	LanguageCodeSpain                                 = "es_ES"
	LanguageCodeItaly                                 = "it_IT"
	LanguageCodeFrance                                = "fr_FR"
	LanguageCodeJapan                                 = "ja_JP"
	LanguageCodeKorea                                 = "ko_KR"
	LanguageCodeMexico                                = "es_MX"
	LanguageCodeArgentina                             = "es_AR"
	LanguageCodeBrazil                                = "pt_BR"
	LanguageCodeUnitedStates                          = "en_US"
	LanguageCodeAustralia                             = "en_AU"
	LanguageCodeRussia                                = "ru_RU"
	LanguageCodeTurkey                                = "tr_TR"
	LanguageCodeMalaysia                              = "ms_MY"
	LanguageCodeRepublicOfThePhilippines              = "en_PH"
	LanguageCodeSingapore                             = "en_SG"
	LanguageCodeThailand                              = "th_TH"
	LanguageCodeVietNam                               = "vn_VN"
	LanguageCodeIndonesia                             = "id_ID"
	LanguageCodeMalaysiaChinese                       = "zh_MY"
	LanguageCodeChina                                 = "zh_CN"
	LanguageCodeTaiwan                                = "zh_TW"
)

// All possible queues
const (
	QueueRankedSolo            queue = "RANKED_SOLO_5x5"
	QueueRankedFlex                  = "RANKED_FLEX_SR"
	QueueRankedTwistedTreeline       = "RANKED_FLEX_TT"
)

// All possible Tiers
const (
	TierIron     tier = "IRON"
	TierBronze        = "BRONZE"
	TierSilver        = "SILVER"
	TierGold          = "GOLD"
	TierPlatinum      = "PLATINUM"
	TierDiamond       = "DIAMOND"
)

// All possible divisions
const (
	DivisionOne   division = "I"
	DivisionTwo            = "II"
	DivisionThree          = "III"
	DivisionFour           = "IV"
)

const (
	identificationName       identification = "name"
	identificationAccountID                 = "account"
	identificationPUUID                     = "puuid"
	identificationSummonerID                = "summonerID"
)

// All existing regions
const (
	RegionBrasil            Region = "br1"
	RegionEuropeNorthEast          = "eun1"
	RegionEuropeWest               = "euw1"
	RegionJapan                    = "jp1"
	RegionKorea                    = "kr"
	RegionLatinAmericaNorth        = "la1"
	RegionLatinAmericaSouth        = "la2"
	RegionNorthAmerica             = "na1"
	RegionOceania                  = "oc1"
	RegionTurkey                   = "tr1"
	RegionRussia                   = "ru"
	RegionPBE                      = "pbe1"
)

const (
	dataDragonBaseURL        dataDragonURL = "ddragon.leagueoflegends.com"
	dataDragonDataURLFormat                = dataDragonBaseURL + "/cdn/%s/data/%s"
	dataDragonImageURLFormat               = dataDragonBaseURL + "/cdn/%s/img"
)

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
	endpointSummonerBase                 = "/lol/summoner/v4"
	endpointGetSummonerBySummonerID      = endpointSummonerBase + "/summoners/%s"
	endpointGetSummonerBy                = endpointSummonerBase + "/summoners/by-%s/%s"
	endpointSpectatorBase                = "/lol/spectator/v4"
	endpointGetCurrentGame               = endpointSpectatorBase + "/active-games/by-summoner/%s"
	endpointGetFeaturedGames             = endpointSpectatorBase + "/featured-games"
)

var (
	// LanguageCodes is a list of all possible language codes
	LanguageCodes = []languageCode{
		LanguageCodeCzechRepublic,
		LanguageCodeGreece,
		LanguageCodePoland,
		LanguageCodeRomania,
		LanguageCodeHungary,
		LanguageCodeUnitedKingdom,
		LanguageCodeGermany,
		LanguageCodeSpain,
		LanguageCodeItaly,
		LanguageCodeFrance,
		LanguageCodeJapan,
		LanguageCodeKorea,
		LanguageCodeMexico,
		LanguageCodeArgentina,
		LanguageCodeBrazil,
		LanguageCodeUnitedStates,
		LanguageCodeAustralia,
		LanguageCodeRussia,
		LanguageCodeTurkey,
		LanguageCodeMalaysia,
		LanguageCodeRepublicOfThePhilippines,
		LanguageCodeSingapore,
		LanguageCodeThailand,
		LanguageCodeVietNam,
		LanguageCodeIndonesia,
		LanguageCodeMalaysiaChinese,
		LanguageCodeChina,
		LanguageCodeTaiwan,
	}

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

	// Regions is a list of all available regions
	Regions = []Region{
		RegionBrasil,
		RegionEuropeNorthEast,
		RegionEuropeWest,
		RegionJapan,
		RegionKorea,
		RegionLatinAmericaNorth,
		RegionLatinAmericaSouth,
		RegionNorthAmerica,
		RegionOceania,
		RegionTurkey,
		RegionRussia,
		RegionPBE,
	}
)
