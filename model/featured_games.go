package model

type FeaturedGames struct {
	ClientRefreshInterval int                `json:"clientRefreshInterval"`
	GameList              []FeaturedGameInfo `json:"gameList"`
}
