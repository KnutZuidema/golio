package model

// Tournament contains the settings of a previously created tournament
type Tournament struct {
	Map          string   `json:"map"`
	Code         string   `json:"code"`
	Spectators   string   `json:"spectators"`
	Region       string   `json:"region"`
	ProviderID   int      `json:"providerId"`
	TeamSize     int      `json:"teamSize"`
	Participants []string `json:"participants"`
	PickType     string   `json:"pickType"`
	TournamentID int      `json:"tournamentId"`
	LobbyName    string   `json:"lobbyName"`
	Password     string   `json:"password"`
	ID           int      `json:"id"`
	MetaData     string   `json:"metaData"`
}
