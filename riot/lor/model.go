package lor

// Player represents a ranked Legends of Runeterra player.
type Player struct {
	Name         string `json:"name"`
	Rank         int    `json:"rank"`
	LeaguePoints int    `json:"lp"`
}
