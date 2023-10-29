package val

import "github.com/KnutZuidema/golio/api"

const (
	endpointBase                  = "/val"
	endpointContentBase           = endpointBase + "/content/v1"
	endPointGetContent            = endpointContentBase + "/contents?%s"
	endpointStatusBase            = endpointBase + "/status/v1"
	endpointGetPlatformData       = endpointStatusBase + "/platform-data"
	endpointRankedBase            = endpointBase + "/ranked/v1"
	endpointGetLeaderboardByActID = endpointRankedBase + "/leaderboards/by-act/%s"
	endpointMatchBase             = endpointBase + "/match/v1"
	endpointMatchByID             = endpointMatchBase + "/matches/%s"
	endpointMatchListByPUUID      = endpointMatchBase + "/matchlists/by-puuid/%s"
	endpointRecentMatchesByQueue  = endpointMatchBase + "/recent-matches/by-queue/%s"
)

// All existing regions
const (
	RegionAsiaPacific  api.Region = "ap"
	RegionBrazil       api.Region = "br"
	RegionESPORTS      api.Region = "esports"
	RegionEurope       api.Region = "eu"
	RegionKorea        api.Region = "kr"
	RegionLatinAmerica api.Region = "latam"
	RegionNorthAmerica api.Region = "na"
)

var (
	// Regions is a list of all available regions
	Regions = []api.Region{
		RegionAsiaPacific,
		RegionBrazil,
		RegionESPORTS,
		RegionEurope,
		RegionKorea,
		RegionLatinAmerica,
		RegionNorthAmerica,
	}

	// RegionToRoute maps each region to its route
	RegionToRoute = map[api.Region]api.Route{
		RegionAsiaPacific:  api.RouteAsia,
		RegionBrazil:       api.RouteAmericas,
		RegionESPORTS:      api.RouteAmericas, //todo: fix the correct route for esports region
		RegionEurope:       api.RouteEurope,
		RegionKorea:        api.RouteAsia,
		RegionLatinAmerica: api.RouteAmericas,
		RegionNorthAmerica: api.RouteAmericas,
	}
)

// Locale string value for language
type Locale string

// All possible values of Locale
const (
	LocaleUnitedArabEmirates Locale = "ar-AE"
	LocaleGermany            Locale = "de-DE"
	LocaleUnitedKingdom      Locale = "en-GB"
	LocaleUnitedStates       Locale = "en-US"
	LocaleSpain              Locale = "es-ES"
	LocaleMexico             Locale = "es-MX"
	LocaleFrance             Locale = "fr-FR"
	LocaleIndonesia          Locale = "id-ID"
	LocaleItaly              Locale = "it-IT"
	LocaleJapan              Locale = "ja-JP"
	LocaleSouthKorea         Locale = "ko-KR"
	LocalePoland             Locale = "pl-PL"
	LocaleBrazil             Locale = "pt-BR"
	LocaleRussia             Locale = "ru-RU"
	LocaleThailand           Locale = "th-TH"
	LocaleTurkish            Locale = "tr-TR"
	LocaleVietnam            Locale = "vi-VN"
	LocaleChina              Locale = "zh-CN"
	LocaleTaiwan             Locale = "zh-TW"
)
