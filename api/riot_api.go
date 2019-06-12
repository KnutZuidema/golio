package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/KnutZuidema/golio/model"

	log "github.com/sirupsen/logrus"
)

// RiotAPIClient provides access to all Riot API endpoints
type RiotAPIClient struct {
	logger log.FieldLogger
	Region Region
	apiKey string
	client Doer
}

// NewRiotAPIClient returns a new api client for the Riot API
func NewRiotAPIClient(region Region, apiKey string, client Doer, logger log.FieldLogger) *RiotAPIClient {
	return &RiotAPIClient{
		Region: region,
		apiKey: apiKey,
		client: client,
		logger: logger.WithField("client", "riot api"),
	}
}

// GetSummonerByName returns the summoner with the given summoner name
func (c RiotAPIClient) GetSummonerByName(name string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationName, name)
}

// GetSummonerByAccount returns the summoner with the given account ID
func (c RiotAPIClient) GetSummonerByAccount(id string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationAccountID, id)
}

// GetSummonerByPUUID returns the summoner with the given PUUID
func (c RiotAPIClient) GetSummonerByPUUID(puuid string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationPUUID, puuid)
}

// GetSummonerBySummonerID returns the summoner with the given ID
func (c RiotAPIClient) GetSummonerBySummonerID(summonerID string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationSummonerID, summonerID)
}

// GetChampionMasteries returns information about masteries for the summoner with the given ID
func (c RiotAPIClient) GetChampionMasteries(summonerID string) ([]*model.ChampionMastery, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChampionMasteries",
		"Region": c.Region,
	})
	var masteries []*model.ChampionMastery
	if err := c.getInto(
		fmt.Sprintf(endpointGetChampionMasteries, summonerID),
		&masteries,
	); err != nil {
		logger.Error(err)
		return nil, err
	}
	return masteries, nil
}

// GetChampionMastery returns information about the mastery of the champion with the given ID the summoner with the
// given ID has
func (c RiotAPIClient) GetChampionMastery(summonerID, championID string) (*model.ChampionMastery, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChampionMastery",
		"Region": c.Region,
	})
	var mastery *model.ChampionMastery
	if err := c.getInto(
		fmt.Sprintf(endpointGetChampionMastery, summonerID, championID),
		&mastery,
	); err != nil {
		logger.Error(err)
		return nil, err
	}
	return mastery, nil
}

// GetChampionMasteryTotalScore returns the accumulated mastery score of all champions played by the summoner with the
// given ID
func (c RiotAPIClient) GetChampionMasteryTotalScore(summonerID string) (int, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChampionMasteryTotalScore",
		"Region": c.Region,
	})
	var score int
	if err := c.getInto(fmt.Sprintf(endpointGetChampionMasteryTotalScore, summonerID), &score); err != nil {
		logger.Error(err)
		return 0, err
	}
	return score, nil
}

// GetFreeChampionRotation returns information about the current free champion rotation
func (c RiotAPIClient) GetFreeChampionRotation() (*model.ChampionInfo, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetFreeChampionRotation",
		"Region": c.Region,
	})
	var info *model.ChampionInfo
	if err := c.getInto(endpointGetFreeChampionRotation, &info); err != nil {
		logger.Error(err)
		return nil, err
	}
	return info, nil
}

// GetChallengerLeague returns the current Challenger league for the Region
func (c RiotAPIClient) GetChallengerLeague(queue queue) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChallengerLeague",
		"Region": c.Region,
	})
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetChallengerLeague, queue), &list); err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

// GetGrandmasterLeague returns the current Grandmaster league for the Region
func (c RiotAPIClient) GetGrandmasterLeague(queue queue) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetGrandmasterLeague",
		"Region": c.Region,
	})
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetGrandmasterLeague, queue), &list); err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

// GetMasterLeague returns the current Master league for the Region
func (c RiotAPIClient) GetMasterLeague(queue queue) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMasterLeague",
		"Region": c.Region,
	})
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetMasterLeague, queue), &list); err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

// GetLeaguesBySummoner returns all leagues a summoner with the given ID is in
func (c RiotAPIClient) GetLeaguesBySummoner(summonerID string) ([]*model.LeagueEntry, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLeaguesBySummoner",
		"Region": c.Region,
	})
	var leagues []*model.LeagueEntry
	if err := c.getInto(fmt.Sprintf(endpointGetLeaguesBySummoner, summonerID), &leagues); err != nil {
		logger.Error(err)
		return nil, err
	}
	return leagues, nil
}

// GetLeagues returns all players with a a league specified by its queue, tier and division
func (c RiotAPIClient) GetLeagues(queue queue, tier tier, division division) ([]*model.LeagueEntry, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLeagues",
		"Region": c.Region,
	})
	var leagues []*model.LeagueEntry
	if err := c.getInto(fmt.Sprintf(endpointGetLeagues, queue, tier, division), &leagues); err != nil {
		logger.Error(err)
		return nil, err
	}
	return leagues, nil
}

// GetLeague returns a ranked league with the specified ID
func (c RiotAPIClient) GetLeague(leagueID string) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLeague",
		"Region": c.Region,
	})
	var leagues *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetLeague, leagueID), &leagues); err != nil {
		logger.Error(err)
		return nil, err
	}
	return leagues, nil
}

// GetStatus returns the current status of the services for the Region
func (c RiotAPIClient) GetStatus() (*model.Status, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetStatus",
		"Region": c.Region,
	})
	var status *model.Status
	if err := c.getInto(endpointGetStatus, &status); err != nil {
		logger.Error(err)
		return nil, err
	}
	return status, nil
}

// GetMatch returns a match specified by its ID
func (c RiotAPIClient) GetMatch(id int) (*model.Match, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMatch",
		"Region": c.Region,
	})
	var match *model.Match
	if err := c.getInto(fmt.Sprintf(endpointGetMatch, id), &match); err != nil {
		logger.Error(err)
		return nil, err
	}
	return match, nil
}

// GetMatchesByAccount returns a specified range of matches played on the account
func (c RiotAPIClient) GetMatchesByAccount(accountID string, beginIndex, endIndex int) (*model.Matchlist, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMatchesByAccount",
		"Region": c.Region,
	})
	var matches *model.Matchlist
	if err := c.getInto(
		fmt.Sprintf(endpointGetMatchesByAccount, accountID, beginIndex, endIndex),
		&matches,
	); err != nil {
		logger.Error(err)
		return nil, err
	}
	return matches, nil
}

// MatchStreamValue value returned by GetMatchesByAccountStream, containing either a reference to a match or an error
type MatchStreamValue struct {
	*model.MatchReference
	Error error
}

// GetMatchesByAccountStream returns all matches played on this account as a stream, requesting new until there are no
// more new games
func (c RiotAPIClient) GetMatchesByAccountStream(accountID string) <-chan MatchStreamValue {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMatchesByAccountStream",
		"Region": c.Region,
	})
	cMatches := make(chan MatchStreamValue, 100)
	go func() {
		start := 0
		for {
			matches, err := c.GetMatchesByAccount(accountID, start, start+100)
			if err != nil {
				logger.Error(err)
				cMatches <- MatchStreamValue{Error: err}
				return
			}
			for _, match := range matches.Matches {
				m := new(model.MatchReference)
				*m = match
				cMatches <- MatchStreamValue{MatchReference: m}
			}
			if len(matches.Matches) < 100 {
				cMatches <- MatchStreamValue{Error: io.EOF}
				return
			}
			start += 100
		}
	}()
	return cMatches
}

// GetCurrentGame returns a currently running game for a summoner
func (c RiotAPIClient) GetCurrentGame(summonerID string) (*model.CurrentGameInfo, error) {
	logger := c.logger.WithFields(log.Fields{
		"method":     "GetCurrentGame",
		"Region":     c.Region,
		"summonerID": summonerID,
	})
	var games model.CurrentGameInfo
	if err := c.getInto(fmt.Sprintf(endpointGetCurrentGame, summonerID), &games); err != nil {
		logger.Error(err)
		return nil, err
	}
	return &games, nil
}

// GetFeaturedGames returns the currently featured games
func (c RiotAPIClient) GetFeaturedGames() (*model.FeaturedGames, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetFeaturedGames",
		"Region": c.Region,
	})
	var games model.FeaturedGames
	if err := c.getInto(endpointGetFeaturedGames, &games); err != nil {
		logger.Error(err)
		return nil, err
	}
	return &games, nil
}

// CreateTournamentCodes creates a specified amount of codes for a tournament.
// For more information about the parameters see the documentation for TournamentCodeParameters.
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (c RiotAPIClient) CreateTournamentCodes(tournamentID, count int, parameters *model.TournamentCodeParameters,
	useStub bool) ([]string, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "CreateTournamentCodes",
		"Region": c.Region,
		"stub":   useStub,
	})
	endpoint := endpointCreateTournamentCodes
	if useStub {
		endpoint = endpointCreateStubTournamentCodes
	}
	var codes []string
	if err := c.postInto(fmt.Sprintf(endpoint, count, tournamentID), parameters, &codes); err != nil {
		logger.Error(err)
		return nil, err
	}
	return codes, nil
}

// GetLobbyEvents returns the lobby events for a lobby specified by the tournament code
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (c RiotAPIClient) GetLobbyEvents(code string, useStub bool) (*model.LobbyEventList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLobbyEvents",
		"Region": c.Region,
		"stub":   useStub,
	})
	endpoint := endpointGetLobbyEvents
	if useStub {
		endpoint = endpointGetStubLobbyEvents
	}
	var events model.LobbyEventList
	if err := c.getInto(fmt.Sprintf(endpoint, code), &events); err != nil {
		logger.Error(err)
		return nil, err
	}
	return &events, nil
}

// CreateTournamentProvider creates a tournament provider and returns the ID.
// For more information about the parameters see the documentation for ProviderRegistrationParameters.
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (c RiotAPIClient) CreateTournamentProvider(parameters *model.ProviderRegistrationParameters,
	useStub bool) (int, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "CreateTournamentProvider",
		"Region": c.Region,
		"stub":   useStub,
	})
	endpoint := endpointCreateTournamentProvider
	if useStub {
		endpoint = endpointCreateStubTournamentProvider
	}
	var id int
	if err := c.postInto(endpoint, parameters, &id); err != nil {
		logger.Error(err)
		return 0, err
	}
	return id, nil
}

// CreateTournament creates a tournament and returns the ID.
// For more information about the parameters see the documentation for TournamentRegistrationParameters.
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (c RiotAPIClient) CreateTournament(parameters *model.TournamentRegistrationParameters, useStub bool) (int, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "CreateTournament",
		"Region": c.Region,
		"stub":   useStub,
	})
	endpoint := endpointCreateTournament
	if useStub {
		endpoint = endpointCreateStubTournament
	}
	var id int
	if err := c.postInto(endpoint, parameters, &id); err != nil {
		logger.Error(err)
		return 0, err
	}
	return id, nil
}

func (c RiotAPIClient) getSummonerBy(by identification, value string) (*model.Summoner, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "getSummonerBy",
		"Region": c.Region,
	})
	var endpoint string
	switch by {
	case identificationSummonerID:
		endpoint = fmt.Sprintf(endpointGetSummonerBySummonerID, value)
	default:
		endpoint = fmt.Sprintf(endpointGetSummonerBy, by, value)
	}
	var summoner *model.Summoner
	if err := c.getInto(endpoint, &summoner); err != nil {
		logger.Error(err)
		return nil, err
	}
	return summoner, nil
}

func (c RiotAPIClient) getInto(endpoint string, target interface{}) error {
	logger := c.logger.WithFields(log.Fields{
		"method":   "getInto",
		"Region":   c.Region,
		"endpoint": endpoint,
	})
	response, err := c.get(endpoint)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (c RiotAPIClient) postInto(endpoint string, body, target interface{}) error {
	logger := c.logger.WithFields(log.Fields{
		"method":   "postInto",
		"Region":   c.Region,
		"endpoint": endpoint,
	})
	response, err := c.post(endpoint, body)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (c RiotAPIClient) get(endpoint string) (*http.Response, error) {
	return c.doRequest("GET", endpoint, nil)
}

func (c RiotAPIClient) post(endpoint string, body interface{}) (*http.Response, error) {
	logger := c.logger.WithFields(log.Fields{
		"method":   "post",
		"Region":   c.Region,
		"endpoint": endpoint,
	})
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		logger.Error(err)
		return nil, err
	}
	return c.doRequest("POST", endpoint, buf)
}

func (c RiotAPIClient) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	logger := c.logger.WithFields(log.Fields{
		"method":   "doRequest",
		"Region":   c.Region,
		"endpoint": endpoint,
	})
	request, err := c.newRequest(method, endpoint, body)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if response.StatusCode == http.StatusServiceUnavailable {
		logger.Info("service unavailable, retrying")
		time.Sleep(time.Second)
		response, err = c.client.Do(request)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}
	if response.StatusCode == http.StatusTooManyRequests {
		retry := response.Header.Get("Retry-After")
		seconds, err := strconv.Atoi(retry)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		logger.Infof("rate limited, waiting %d seconds", seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.doRequest(method, endpoint, body)
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		logger.Errorf("error response: %v", response.Status)
		err, ok := StatusToError[response.StatusCode]
		if !ok {
			err = Error{
				Message:    "unknown error reason",
				StatusCode: response.StatusCode,
			}
		}
		return nil, err
	}
	return response, nil
}

func (c RiotAPIClient) newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "newRequest",
		"Region": c.Region,
	})
	request, err := http.NewRequest(method, fmt.Sprintf(apiURLFormat, scheme, c.Region, baseURL, endpoint), body)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	request.Header.Add(apiTokenHeaderKey, c.apiKey)
	request.Header.Add("Accept", "application/json")
	return request, nil
}
