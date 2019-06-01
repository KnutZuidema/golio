package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/pkg/model"
)

type RiotAPIClient struct {
	logger log.FieldLogger
	Region region
	apiKey string
	client *http.Client
}

func NewRiotAPIClient(region region, apiKey string, client *http.Client, logger log.FieldLogger) *RiotAPIClient {
	return &RiotAPIClient{
		Region: region,
		apiKey: apiKey,
		client: client,
		logger: logger.WithField("client", "riot api"),
	}
}

func (c RiotAPIClient) GetSummonerByName(name string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationName, name)
}

func (c RiotAPIClient) GetSummonerByAccount(id string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationAccountID, id)
}

func (c RiotAPIClient) GetSummonerByPUUID(puuid string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationPUUID, puuid)
}

func (c RiotAPIClient) GetSummonerBySummonerID(summonerID string) (*model.Summoner, error) {
	return c.getSummonerBy(identificationSummonerID, summonerID)
}

func (c RiotAPIClient) GetChampionMasteries(summonerID string) ([]*model.ChampionMastery, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChampionMasteries",
		"region": c.Region,
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

func (c RiotAPIClient) GetChampionMastery(summonerID string, championID string) (*model.ChampionMastery, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChampionMastery",
		"region": c.Region,
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

func (c RiotAPIClient) GetChampionMasteryTotalScore(summonerID string) (int, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChampionMasteryTotalScore",
		"region": c.Region,
	})
	response, err := c.get(fmt.Sprintf(endpointGetChampionMasteryTotalScore, summonerID))
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(response.Body)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	score, err := strconv.Atoi(string(buffer.Bytes()))
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	return score, nil
}

func (c RiotAPIClient) GetFreeChampionRotation() (*model.ChampionInfo, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetFreeChampionRotation",
		"region": c.Region,
	})
	var info *model.ChampionInfo
	if err := c.getInto(endpointGetFreeChampionRotation, &info); err != nil {
		logger.Error(err)
		return nil, err
	}
	return info, nil
}

func (c RiotAPIClient) GetChallengerLeague(queue queue) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetChallengerLeague",
		"region": c.Region,
	})
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetChallengerLeague, queue), &list); err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

func (c RiotAPIClient) GetGrandmasterLeague(queue queue) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetGrandmasterLeague",
		"region": c.Region,
	})
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetGrandmasterLeague, queue), &list); err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

func (c RiotAPIClient) GetMasterLeague(queue queue) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMasterLeague",
		"region": c.Region,
	})
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetMasterLeague, queue), &list); err != nil {
		logger.Error(err)
		return nil, err
	}
	return list, nil
}

func (c RiotAPIClient) GetLeaguesBySummoner(summonerID string) ([]*model.LeagueEntry, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLeaguesBySummoner",
		"region": c.Region,
	})
	var leagues []*model.LeagueEntry
	if err := c.getInto(fmt.Sprintf(endpointGetLeaguesBySummoner, summonerID), &leagues); err != nil {
		logger.Error(err)
		return nil, err
	}
	return leagues, nil
}

func (c RiotAPIClient) GetLeagues(queue queue, tier tier, division division) ([]*model.LeagueEntry, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLeagues",
		"region": c.Region,
	})
	var leagues []*model.LeagueEntry
	if err := c.getInto(fmt.Sprintf(endpointGetLeagues, queue, tier, division), &leagues); err != nil {
		logger.Error(err)
		return nil, err
	}
	return leagues, nil
}

func (c RiotAPIClient) GetLeague(leagueID string) (*model.LeagueList, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetLeague",
		"region": c.Region,
	})
	var leagues *model.LeagueList
	if err := c.getInto(fmt.Sprintf(endpointGetLeague, leagueID), &leagues); err != nil {
		logger.Error(err)
		return nil, err
	}
	return leagues, nil
}

func (c RiotAPIClient) GetStatus() (*model.Status, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetStatus",
		"region": c.Region,
	})
	var status *model.Status
	if err := c.getInto(endpointGetStatus, &status); err != nil {
		logger.Error(err)
		return nil, err
	}
	return status, nil
}

func (c RiotAPIClient) GetMatch(id int) (*model.Match, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMatch",
		"region": c.Region,
	})
	var match *model.Match
	if err := c.getInto(fmt.Sprintf(endpointGetMatch, id), &match); err != nil {
		logger.Error(err)
		return nil, err
	}
	return match, nil
}

func (c RiotAPIClient) GetMatchesByAccount(accountID string, beginIndex, endIndex int) (*model.Matchlist, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMatchesByAccount",
		"region": c.Region,
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

func (c RiotAPIClient) GetMatchesByAccountStream(accountID string) <-chan struct {
	*model.MatchReference
	error
} {
	logger := c.logger.WithFields(log.Fields{
		"method": "GetMatchesByAccountStream",
		"region": c.Region,
	})
	cMatches := make(chan struct {
		*model.MatchReference
		error
	}, 100)
	go func() {
		start := 0
		for {
			matches, err := c.GetMatchesByAccount(accountID, start, start+100)
			if err != nil {
				logger.Error(err)
				cMatches <- struct {
					*model.MatchReference
					error
				}{error: err}
				return
			}
			for _, match := range matches.Matches {
				m := new(model.MatchReference)
				*m = match
				logger.Infof("streaming match %v", match.GameID)
				cMatches <- struct {
					*model.MatchReference
					error
				}{MatchReference: m}
			}
			if len(matches.Matches) < 100 {
				cMatches <- struct {
					*model.MatchReference
					error
				}{error: io.EOF}
				return
			}
			start += 100
		}
	}()
	return cMatches
}

func (c RiotAPIClient) getSummonerBy(by identification, value string) (*model.Summoner, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "getSummonerBy",
		"region": c.Region,
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
		"region":   c.Region,
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

func (c RiotAPIClient) get(endpoint string) (*http.Response, error) {
	return c.doRequest("GET", endpoint, "")
}

func (c RiotAPIClient) doRequest(method, endpoint, body string) (*http.Response, error) {
	logger := c.logger.WithFields(log.Fields{
		"method":   "doRequest",
		"region":   c.Region,
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
		return nil, fmt.Errorf("error response: %v", response.Status)
	}
	return response, nil
}

func (c RiotAPIClient) newRequest(method, endpoint, body string) (*http.Request, error) {
	logger := c.logger.WithFields(log.Fields{
		"method": "newRequest",
		"region": c.Region,
	})
	request, err := http.NewRequest(method, fmt.Sprintf(apiURLFormat, scheme, c.Region, baseURL, endpoint),
		strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	request.Header.Add(apiTokenHeaderKey, c.apiKey)
	request.Header.Add("Accept", "application/json")
	return request, nil
}
