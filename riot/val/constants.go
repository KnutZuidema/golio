package val

import "github.com/yigithanbalci/golio/api"

const (
	endpointBase                  = "/val"
	endpointContentBase           = endpointBase + "/content/v1"
	endPointGetContent            = endpointContentBase + "/contents?%s"
	endpointStatusBase            = endpointBase + "/status/v1"
	endpointGetPlatformData       = endpointStatusBase + "/platform-data"
	endpointRankedBase            = endpointBase + "/ranked/v1"
	endpointGetLeaderboardByActId = endpointRankedBase + "/leaderboards/by-act/%s"
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

type LocalizedNamesDto string

const (
	AR_AE LocalizedNamesDto = "ar-AE"
	DE_DE LocalizedNamesDto = "de-DE"
	EN_GB LocalizedNamesDto = "en-GB"
	EN_US LocalizedNamesDto = "en-US"
	ES_ES LocalizedNamesDto = "es-ES"
	ES_MX LocalizedNamesDto = "es-MX"
	FR_FR LocalizedNamesDto = "fr-FR"
	ID_ID LocalizedNamesDto = "id-ID"
	IT_IT LocalizedNamesDto = "it-IT"
	JA_JP LocalizedNamesDto = "ja-JP"
	KO_KR LocalizedNamesDto = "ko-KR"
	PL_PL LocalizedNamesDto = "pl-PL"
	PT_BR LocalizedNamesDto = "pt-BR"
	RU_RU LocalizedNamesDto = "ru-RU"
	TH_TH LocalizedNamesDto = "th-TH"
	TR_TR LocalizedNamesDto = "tr-TR"
	VI_VN LocalizedNamesDto = "vi-VN"
	ZH_CN LocalizedNamesDto = "zh-CN"
	ZH_TW LocalizedNamesDto = "zh-TW"
)
