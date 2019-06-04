package model

// FeaturedGames represents a list of featured games
type FeaturedGames struct {
	ClientRefreshInterval int                `json:"clientRefreshInterval"`
	GameList              []FeaturedGameInfo `json:"gameList"`
}
