package model

// TournamentRegistrationParameters parameters required for creating a tournament
type TournamentRegistrationParameters struct {
	// The provider ID to specify the regional registered provider data to associate this tournament.
	ProviderID int `json:"providerId"`
	// The optional name of the tournament.
	Name string `json:"name"`
}
