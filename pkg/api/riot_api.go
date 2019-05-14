package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/KnutZuidema/riot-api-wrapper/pkg/model"
)

type RiotAPIClient struct {
	region region
	apiKey string
	client *http.Client
}

func NewRiotAPIClient(region region, apiKey string, client *http.Client) *RiotAPIClient {
	return &RiotAPIClient{region: region, apiKey: apiKey, client: client}
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
	var masteries []*model.ChampionMastery
	if err := c.getInto(
		fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-summoner/%s", summonerID),
		&masteries,
	); err != nil {
		return nil, err
	}
	return masteries, nil
}

func (c RiotAPIClient) GetChampionMastery(summonerID string, championID int) (*model.ChampionMastery, error) {
	var mastery *model.ChampionMastery
	if err := c.getInto(
		fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-summoner/%s/by-champion/%d",
			summonerID,
			championID),
		&mastery,
	); err != nil {
		return nil, err
	}
	return mastery, nil
}

func (c RiotAPIClient) GetChampionMasteryTotalScore(summonerID string) (int, error) {
	response, err := c.get(fmt.Sprintf("/lol/champion-mastery/v4/scores/by-summoner/%s", summonerID))
	if err != nil {
		return 0, err
	}
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(response.Body)
	if err != nil {
		return 0, err
	}
	score, err := strconv.Atoi(string(buffer.Bytes()))
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (c RiotAPIClient) GetFreeChampionRotation() (*model.ChampionInfo, error) {
	var info *model.ChampionInfo
	if err := c.getInto("/lol/platform/v3/champion-rotations", &info); err != nil {
		return nil, err
	}
	return info, nil
}

func (c RiotAPIClient) GetChallengerLeague(queue queue) (*model.LeagueList, error) {
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf("/lol/league/v4/challengerleagues/by-queue/%s", queue), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c RiotAPIClient) GetGrandmasterLeague(queue queue) (*model.LeagueList, error) {
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf("/lol/league/v4/grandmasterleagues/by-queue/%s", queue), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c RiotAPIClient) GetMasterLeague(queue queue) (*model.LeagueList, error) {
	var list *model.LeagueList
	if err := c.getInto(fmt.Sprintf("/lol/league/v4/masterleagues/by-queue/%s", queue), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c RiotAPIClient) GetSummonerLeagues(summonerID string) ([]*model.LeagueEntry, error) {
	var leagues []*model.LeagueEntry
	if err := c.getInto(fmt.Sprintf("/lol/league/v4/entries/by-summoner/%s", summonerID), &leagues); err != nil {
		return nil, err
	}
	return leagues, nil
}

func (c RiotAPIClient) GetLeagues(queue queue, tier tier, division division) ([]*model.LeagueEntry, error) {
	var leagues []*model.LeagueEntry
	if err := c.getInto(fmt.Sprintf("/lol/league/v4/entries/%s/%s/%s", queue, tier, division), &leagues); err != nil {
		return nil, err
	}
	return leagues, nil
}

func (c RiotAPIClient) GetLeague(leagueID string) (*model.LeagueList, error) {
	var leagues *model.LeagueList
	if err := c.getInto(fmt.Sprintf("/lol/league/v4/leagues/%s", leagueID), &leagues); err != nil {
		return nil, err
	}
	return leagues, nil
}

func (c RiotAPIClient) GetStatus() (*model.Status, error) {
	var status *model.Status
	if err := c.getInto("/lol/status/v3/shard-data", &status); err != nil {
		return nil, err
	}
	return status, nil
}

func (c RiotAPIClient) GetMatch(id string) (*model.Match, error) {
	var match *model.Match
	if err := c.getInto(fmt.Sprintf("/lol/match/v4/matches/%s", id), &match); err != nil {
		return nil, err
	}
	return match, nil
}

func (c RiotAPIClient) GetMatchesByAccount(accountID string, beginIndex, endIndex int) (*model.Matchlist, error) {
	var matches *model.Matchlist
	if err := c.getInto(
		fmt.Sprintf("/lol/match/v4/matchlists/by-account/%s?beginIndex=%d&endIndex=%d",
			accountID, beginIndex, endIndex),
		&matches,
	); err != nil {
		return nil, err
	}
	return matches, nil
}

func (c RiotAPIClient) GetMatchesByAccountStream(accountID string) <-chan struct {
	*model.MatchReference
	error
} {
	cMatches := make(chan struct {
		*model.MatchReference
		error
	}, 100)
	go func() {
		start := 0
		for {
			matches, err := c.GetMatchesByAccount(accountID, start, start+100)
			if err != nil {
				cMatches <- struct {
					*model.MatchReference
					error
				}{error: err}
				return
			}
			for _, match := range matches.Matches {
				m := new(model.MatchReference)
				*m = match
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
	var endpoint string
	switch by {
	case identificationSummonerID:
		endpoint = fmt.Sprintf("/lol/summoner/v4/summoners/%s", value)
	default:
		endpoint = fmt.Sprintf("/lol/summoner/v4/summoners/by-%s/%s", by, value)
	}
	var summoner *model.Summoner
	if err := c.getInto(endpoint, &summoner); err != nil {
		return nil, err
	}
	return summoner, nil
}

func (c RiotAPIClient) getInto(endpoint string, target interface{}) error {
	response, err := c.get(endpoint)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		return err
	}
	return nil
}

func (c RiotAPIClient) get(endpoint string) (*http.Response, error) {
	return c.doRequest("GET", endpoint, "")
}

func (c RiotAPIClient) doRequest(method, endpoint, body string) (*http.Response, error) {
	request, err := c.newRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("error response: %v", response.Status)
	}
	return response, nil
}

func (c RiotAPIClient) newRequest(method, endpoint, body string) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf("https://%s.%s%s", c.region, baseURL, endpoint),
		strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Add(apiTokenHeaderKey, c.apiKey)
	request.Header.Add("Accept", "application/json")
	return request, nil
}
