package model

// ProviderRegistrationParameters parameters required for registering a provider with tournaments for a region
type ProviderRegistrationParameters struct {
	// The provider's callback URL to which tournament game results in this region should be posted. The URL must be
	// well-formed, use the http or https protocol, and use the default port for the protocol (http URLs must use port
	// 80, https URLs must use port 443).
	URL string `json:"url"`
	// The region in which the provider will be running tournaments.
	// (Legal values: BR, EUNE, EUW, JP, LAN, LAS, NA, OCE, PBE, RU, TR)
	Region string `json:"region"`
}
