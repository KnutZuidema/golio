package golio

import (
	"fmt"
	"strings"
)

type identification string
type region string
type queue string
type tier string
type division string
type dataDragonURL string
type languageCode string

const (
	LanguageCodeCzechRepublic              languageCode   = "cs_CZ"
	LanguageCodeGreece                                    = "el_GR"
	LanguageCodePoland                                    = "pl_PL"
	LanguageCodeRomania                                   = "ro_RO"
	LanguageCodeHungary                                   = "hu_HU"
	LanguageCodeUnitedKingdom                             = "en_GB"
	LanguageCodeGermany                                   = "de_DE"
	LanguageCodeSpain                                     = "es_ES"
	LanguageCodeItaly                                     = "it_IT"
	LanguageCodeFrance                                    = "fr_FR"
	LanguageCodeJapan                                     = "ja_JP"
	LanguageCodeKorea                                     = "ko_KR"
	LanguageCodeMexico                                    = "es_MX"
	LanguageCodeArgentina                                 = "es_AR"
	LanguageCodeBrazil                                    = "pt_BR"
	LanguageCodeUnitedStates                              = "en_US"
	LanguageCodeAustralia                                 = "en_AU"
	LanguageCodeRussia                                    = "ru_RU"
	LanguageCodeTurkey                                    = "tr_TR"
	LanguageCodeMalaysia                                  = "ms_MY"
	LanguageCodeRepublicOfThePhilippines                  = "en_PH"
	LanguageCodeSingapore                                 = "en_SG"
	LanguageCodeThailand                                  = "th_TH"
	LanguageCodeVietNam                                   = "vn_VN"
	LanguageCodeIndonesia                                 = "id_ID"
	LanguageCodeMalaysiaChinese                           = "zh_MY"
	LanguageCodeChina                                     = "zh_CN"
	LanguageCodeTaiwan                                    = "zh_TW"
	QueueRankedSolo                        queue          = "RANKED_SOLO_5x5"
	QueueRankedFlex                                       = "RANKED_FLEX_SR"
	QueueRankedTwistedTreeline                            = "RANKED_FLEX_TT"
	TierIron                               tier           = "IRON"
	TierBronze                                            = "BRONZE"
	TierSilver                                            = "SILVER"
	TierGold                                              = "GOLD"
	TierPlatinum                                          = "PLATINUM"
	TierDiamond                                           = "DIAMOND"
	DivisionOne                            division       = "I"
	DivisionTwo                                           = "II"
	DivisionThree                                         = "III"
	DivisionFour                                          = "IV"
	identificationName                     identification = "name"
	identificationAccountID                               = "account"
	identificationPUUID                                   = "puuid"
	identificationSummonerID                              = "summonerID"
	RegionBrasil                           region         = "br1"
	RegionEuropeNorthEast                                 = "eun1"
	RegionEuropeWest                                      = "euw1"
	RegionJapan                                           = "jp1"
	RegionKorea                                           = "kr"
	RegionLatinAmericaNorth                               = "la1"
	RegionLatinAmericaSouth                               = "la2"
	RegionNorthAmerica                                    = "na1"
	RegionOceania                                         = "oc1"
	RegionTurkey                                          = "tr1"
	RegionRussia                                          = "ru"
	RegionPBE                                             = "pbe1"
	dataDragonBaseURL                      dataDragonURL  = "ddragon.leagueoflegends.com"
	dataDragonDataURLFormat                               = dataDragonBaseURL + "/cdn/%s/data/%s"
	dataDragonImageURLFormat                              = dataDragonBaseURL + "/cdn/%s/img"
	dataDragonImageURLFormatWithoutVersion                = dataDragonBaseURL + "/cdn/img"
	apiURLFormat                           string         = "%s://%s.%s%s"
	baseURL                                               = "api.riotgames.com"
	scheme                                                = "https"
	apiTokenHeaderKey                                     = "X-Riot-Token"
	endpointChampionMasteryBase                           = "/lol/champion-mastery/v4"
	endpointGetChampionMasteries                          = endpointChampionMasteryBase + "/champion-masteries/by-summoner/%s"
	endpointGetChampionMastery                            = endpointChampionMasteryBase + "/champion-masteries/by-summoner/%s/by-champion/%s"
	endpointGetChampionMasteryTotalScore                  = endpointChampionMasteryBase + "/scores/by-summoner/%s"
	endpointPlatformBase                                  = "/lol/platform/v3"
	endpointGetFreeChampionRotation                       = endpointPlatformBase + "/champion-rotations"
	endpointLeagueBase                                    = "/lol/league/v4"
	endpointGetChallengerLeague                           = endpointLeagueBase + "/challengerleagues/by-queue/%s"
	endpointGetGrandmasterLeague                          = endpointLeagueBase + "/grandmasterleagues/by-queue/%s"
	endpointGetMasterLeague                               = endpointLeagueBase + "/masterleagues/by-queue/%s"
	endpointGetLeaguesBySummoner                          = endpointLeagueBase + "/entries/by-summoner/%s"
	endpointGetLeagues                                    = endpointLeagueBase + "/entries/%s/%s/%s"
	endpointGetLeague                                     = endpointLeagueBase + "/leagues/%s"
	endpointStatusBase                                    = "/lol/status/v3"
	endpointGetStatus                                     = endpointStatusBase + "/shard-data"
	endpointMatchBase                                     = "/lol/match/v4"
	endpointGetMatch                                      = endpointMatchBase + "/matches/%d"
	endpointGetMatchesByAccount                           = endpointMatchBase + "/matchlists/by-account/%s?beginIndex=%d&endIndex=%d"
	endpointSummonerBase                                  = "/lol/summoner/v4"
	endpointGetSummonerBySummonerID                       = endpointSummonerBase + "/summoners/%s"
	endpointGetSummonerBy                                 = endpointSummonerBase + "/summoners/by-%s/%s"
)

var (
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
	Queues = []queue{
		QueueRankedSolo,
		QueueRankedFlex,
		QueueRankedTwistedTreeline,
	}
	Tiers = []tier{
		TierIron,
		TierBronze,
		TierSilver,
		TierGold,
		TierPlatinum,
		TierDiamond,
	}
	Divisions = []division{
		DivisionOne,
		DivisionTwo,
		DivisionThree,
		DivisionFour,
	}
	Regions = []region{
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
	debugRegionIndex = 0
)

func ParseRegion(r string) (region, error) {
	m := map[string]region{
		"brasil":              RegionBrasil,
		"br":                  RegionBrasil,
		"europe_ne":           RegionEuropeNorthEast,
		"eune":                RegionEuropeNorthEast,
		"europe_w":            RegionEuropeWest,
		"euw":                 RegionEuropeWest,
		"japan":               RegionJapan,
		"jp":                  RegionJapan,
		"korea":               RegionKorea,
		"kr":                  RegionKorea,
		"latin_america_north": RegionLatinAmericaNorth,
		"lan":                 RegionLatinAmericaNorth,
		"latin_america_south": RegionLatinAmericaSouth,
		"las":                 RegionLatinAmericaSouth,
		"north_america":       RegionNorthAmerica,
		"na":                  RegionNorthAmerica,
		"oceania":             RegionOceania,
		"oce":                 RegionOceania,
		"turkey":              RegionTurkey,
		"tr":                  RegionTurkey,
		"russia":              RegionRussia,
		"ru":                  RegionRussia,
		"pbe":                 RegionPBE,
	}
	reg, ok := m[strings.ToLower(r)]
	if !ok {
		return "", fmt.Errorf("could not parse value as region: %s", r)
	}
	return reg, nil
}
