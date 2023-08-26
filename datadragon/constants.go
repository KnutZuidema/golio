package datadragon

type dataDragonURL string

const (
	dataDragonBaseURL        dataDragonURL = "ddragon.leagueoflegends.com"
	dataDragonDataURLFormat                = dataDragonBaseURL + "/cdn/%s/data/%s"
	dataDragonImageURLFormat               = dataDragonBaseURL + "/cdn/%s/img"
)

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
	LanguageCodeVietnam                               = "vi_VN"
	LanguageCodeIndonesia                             = "id_ID"
	LanguageCodeMalaysiaChinese                       = "zh_MY"
	LanguageCodeChina                                 = "zh_CN"
	LanguageCodeTaiwan                                = "zh_TW"
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
		LanguageCodeVietnam,
		LanguageCodeIndonesia,
		LanguageCodeMalaysiaChinese,
		LanguageCodeChina,
		LanguageCodeTaiwan,
	}
)
