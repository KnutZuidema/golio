package api

// Region represents a server region
type Region string

// All existing regions
const (
	RegionBrasil            Region = "br1"
	RegionEuropeNorthEast   Region = "eun1"
	RegionEuropeWest        Region = "euw1"
	RegionJapan             Region = "jp1"
	RegionKorea             Region = "kr"
	RegionLatinAmericaNorth Region = "la1"
	RegionLatinAmericaSouth Region = "la2"
	RegionMiddleEast        Region = "me1"
	RegionNorthAmerica      Region = "na1"
	RegionOceania           Region = "oc1"
	RegionPBE               Region = "pbe1"
	// Deprecated: Use api.RegionSouthEastAsia instead. PH2 got merged into the SEA server on 8th of Jan, 2025.
	RegionPhilippines Region = "sg2" // PH2 got merged into SG2 on Jan. 8th, 2025. Replaced for backwards compatability.
	RegionRussia      Region = "ru"
	// Deprecated: Use api.RegionSouthEastAsia instead. SG2 is now called SEA
	RegionSingapore     Region = "sg2" // SG2 is now called SEA while still running on SG2.
	RegionSouthEastAsia Region = "sg2"
	// Deprecated: Use api.RegionSouthEastAsia instead. TH2 got merged into the SEA server on 8th of Jan, 2025.
	RegionThailand Region = "sg2" // TH2 got merged into SG2 on Jan. 8th, 2025. Replaced for backwards compatability.
	RegionTurkey   Region = "tr1"
	RegionTaiwan   Region = "tw2"
	RegionVietnam  Region = "vn2"
)

// Route represents a server region's route
type Route string

// All existing routes
const (
	RouteAmericas Route = "americas"
	RouteAsia     Route = "asia"
	RouteEurope   Route = "europe"
	RouteSEA      Route = "sea"
)

var (
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
		RegionMiddleEast,
		RegionOceania,
		RegionPBE,
		RegionRussia,
		RegionSouthEastAsia,
		RegionTurkey,
		RegionTaiwan,
		RegionVietnam,
	}

	// RegionToRoute maps each region to its route
	RegionToRoute = map[Region]Route{
		RegionBrasil:            RouteAmericas,
		RegionEuropeNorthEast:   RouteEurope,
		RegionEuropeWest:        RouteEurope,
		RegionJapan:             RouteAsia,
		RegionKorea:             RouteAsia,
		RegionLatinAmericaNorth: RouteAmericas,
		RegionLatinAmericaSouth: RouteAmericas,
		RegionMiddleEast:        RouteEurope,
		RegionNorthAmerica:      RouteAmericas,
		RegionOceania:           RouteSEA,
		RegionRussia:            RouteEurope,
		RegionSouthEastAsia:     RouteSEA,
		RegionTurkey:            RouteEurope,
		RegionTaiwan:            RouteSEA,
		RegionVietnam:           RouteSEA,
	}
)
