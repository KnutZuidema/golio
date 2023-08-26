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
	RegionNorthAmerica      Region = "na1"
	RegionOceania           Region = "oc1"
	RegionPhilippines       Region = "ph2"
	RegionRussia            Region = "ru"
	RegionSingapore         Region = "sg2"
	RegionThailand          Region = "th2"
	RegionTurkey            Region = "tr1"
	RegionTaiwan            Region = "tw2"
	RegionVietnam           Region = "vn2"
	RegionPBE               Region = "pbe1"
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
		RegionOceania,
		RegionPhilippines,
		RegionRussia,
		RegionSingapore,
		RegionThailand,
		RegionTurkey,
		RegionTaiwan,
		RegionVietnam,
		RegionPBE,
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
		RegionNorthAmerica:      RouteAmericas,
		RegionOceania:           RouteAmericas,
		RegionPhilippines:       RouteSEA,
		RegionRussia:            RouteEurope,
		RegionSingapore:         RouteSEA,
		RegionThailand:          RouteSEA,
		RegionTurkey:            RouteEurope,
		RegionTaiwan:            RouteSEA,
		RegionVietnam:           RouteSEA,
	}
)
