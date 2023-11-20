package account

const (
	endpointBase         = "/riot"
	endpointAccountBase  = endpointBase + "/account/v1"
	endpointAccountsBase = endpointAccountBase + "/accounts"
	endpointGetByPuuid   = endpointAccountsBase + "/by-puuid/%s"
	endpointGetByRiotId  = endpointAccountBase + "/by-riot-id/%s/%s"
)
