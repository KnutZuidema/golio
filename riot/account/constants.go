package account

const (
	endpointBase         = "/riot"
	endpointAccountBase  = endpointBase + "/account/v1"
	endpointAccountsBase = endpointAccountBase + "/accounts"
	endpointGetByPUUID   = endpointAccountsBase + "/by-puuid/%s"
	endpointGetByRiotID  = endpointAccountsBase + "/by-riot-id/%s/%s"
)
