package model

// TeamStats represents the stats of a team for a single game
type TeamStats struct {
	// Flag indicating whether or not the team scored the first Dragon kill.
	FirstDragon bool `json:"firstDragon"`
	// Flag indicating whether or not the team destroyed the first inhibitor.
	FirstInhibitor bool `json:"firstInhibitor"`
	// If match queueID has a draft, contains banned champion data, otherwise empty.
	Bans []TeamBan `json:"bans"`
	// Number of times the team killed Baron.
	BaronKills int `json:"baronKills"`
	// Flag indicating whether or not the team scored the first Rift Herald kill.
	FirstRiftHerald bool `json:"firstRiftHerald"`
	// Flag indicating whether or not the team scored the first Baron kill.
	FirstBaron bool `json:"firstBaron"`
	// Number of times the team killed Rift Herald.
	RiftHeraldKills int `json:"riftHeraldKills"`
	// Flag indicating whether or not the team scored the first blood.
	FirstBlood bool `json:"firstBlood"`
	// 100 for blue side. 200 for red side.
	TeamID int `json:"teamId"`
	// Flag indicating whether or not the team destroyed the first tower.
	FirstTower bool `json:"firstTower"`
	// Number of times the team killed Vilemaw.
	VilemawKills int `json:"vilemawKills"`
	// Number of inhibitors the team destroyed.
	InhibitorKills int `json:"inhibitorKills"`
	// Number of towers the team destroyed.
	TowerKills int `json:"towerKills"`
	// For Dominion matches, specifies the points the team had at game end.
	DominionVictoryScore int `json:"dominionVictoryScore"`
	// string indicating whether or not the team won. There are only two values visibile in public match history.
	// (Legal values: Fail, Win)
	Win string `json:"win"`
	// Number of times the team killed Dragon.
	DragonKills int `json:"dragonKills"`
}
