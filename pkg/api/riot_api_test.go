package api

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/KnutZuidema/riot-api-wrapper/pkg/model"
	"github.com/stretchr/testify/assert"
)

const (
	apiKey       = "RGAPI-b52c974c-1284-4dd9-82d6-f6c4d5f405d7"
	summonerName = "InMM BlackScorp"
)

var (
	client   = NewClient(RegionEuropeWest, apiKey, http.DefaultClient)
	summoner *model.Summoner
)

func init() {
	var err error
	summoner, err = client.GetSummonerByName(summonerName)
	if err != nil {
		panic(err)
	}
}

func TestRiotAPIClient_GetSummonerByName(t *testing.T) {
	tests := []struct {
		name         string
		summonerName string
		wantErr      bool
	}{
		{
			name:         "get summoner",
			summonerName: summonerName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			summoner, err = client.GetSummonerByName(tt.summonerName)
			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error is not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.NotNil(t, summoner)
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerByAccount(t *testing.T) {
	tests := []struct {
		name      string
		accountID string
		want      *model.Summoner
		wantErr   bool
	}{
		{
			name:      "get summoner",
			accountID: summoner.AccountID,
			want:      summoner,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetSummonerByAccount(tt.accountID)
			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error is not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.Equal(t, got, tt.want, "responses does not match")
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerByPUUID(t *testing.T) {
	tests := []struct {
		name    string
		puuid   string
		want    *model.Summoner
		wantErr bool
	}{
		{
			name:  "get summoner",
			puuid: summoner.PUUID,
			want:  summoner,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetSummonerByPUUID(tt.puuid)
			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error is not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.Equal(t, got, tt.want, "responses does not match")
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerBySummonerID(t *testing.T) {
	tests := []struct {
		name       string
		summonerID string
		want       *model.Summoner
		wantErr    bool
	}{
		{
			name:       "get summoner",
			summonerID: summoner.ID,
			want:       summoner,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetSummonerBySummonerID(tt.summonerID)
			require.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error is not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.Equal(t, got, tt.want, "responses does not match")
			}
		})
	}
}

func TestRiotAPIClient_GetChampionMasteries(t *testing.T) {
	tests := []struct {
		name       string
		summonerID string
		wantErr    bool
	}{
		{
			name:       "get champion masteries",
			summonerID: summoner.ID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masteries, err := client.GetChampionMasteries(tt.summonerID)
			require.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error was not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.NotNil(t, masteries)
			}
		})
	}
}

func TestRiotAPIClient_GetChampionMastery(t *testing.T) {
	champion, err := client.GetChampion("Anivia")
	require.Nil(t, err)
	championID, err := strconv.Atoi(champion.Key)
	require.Nil(t, err)
	tests := []struct {
		name       string
		summonerID string
		championID int
		wantErr    bool
	}{
		{
			name:       "get champion mastery",
			summonerID: summoner.ID,
			championID: championID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mastery, err := client.GetChampionMastery(tt.summonerID, tt.championID)
			require.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error was not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.NotNil(t, mastery)
			}
		})
	}
}

func TestRiotAPIClient_GetChampionMasteryTotalScore(t *testing.T) {
	tests := []struct {
		name       string
		summonerID string
		wantErr    bool
	}{
		{
			name:       "get total mastery score",
			summonerID: summoner.ID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score, err := client.GetChampionMasteryTotalScore(tt.summonerID)
			require.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error was not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.NotNil(t, score)
			}
		})
	}
}

func TestRiotAPIClient_GetFreeChampionRotation(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "get rotation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetFreeChampionRotation()
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetChallengerLeague(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		queue   queue
	}{
		{
			name:  "get solo challenger queue",
			queue: QueueRankedSolo,
		},
		{
			name:  "get flex challenger queue",
			queue: QueueRankedFlex,
		},
		{
			name:  "get twisted treeline challenger queue",
			queue: QueueRankedTwistedTreeline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetChallengerLeague(tt.queue)
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetGrandmasterLeague(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		queue   queue
	}{
		{
			name:  "get solo grandmaster queue",
			queue: QueueRankedSolo,
		},
		{
			name:  "get flex grandmaster queue",
			queue: QueueRankedFlex,
		},
		{
			name:  "get twisted treeline grandmaster queue",
			queue: QueueRankedTwistedTreeline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetGrandmasterLeague(tt.queue)
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetMasterLeague(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		queue   queue
	}{
		{
			name:  "get solo master queue",
			queue: QueueRankedSolo,
		},
		{
			name:  "get flex master queue",
			queue: QueueRankedFlex,
		},
		{
			name:  "get twisted treeline master queue",
			queue: QueueRankedTwistedTreeline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetMasterLeague(tt.queue)
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerLeagues(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    bool
		summonerID string
	}{
		{
			name:       "get solo queue",
			summonerID: summoner.ID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetLeaguesBySummoner(tt.summonerID)
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetLeagues(t *testing.T) {
	type test struct {
		name     string
		wantErr  bool
		queue    queue
		tier     tier
		division division
	}
	tests := []test{}
	for _, q := range Queues {
		for _, t := range Tiers {
			for _, d := range Divisions {
				tests = append(tests, test{
					name:     fmt.Sprintf("get leagues %s %s %s", q, t, d),
					queue:    q,
					tier:     t,
					division: d,
				})
			}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetLeagues(tt.queue, tt.tier, tt.division)
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetLeague(t *testing.T) {
	leagues, err := client.GetChallengerLeague(QueueRankedSolo)
	require.Nil(t, err)
	require.NotNil(t, leagues)
	leagueID := leagues.LeagueID
	tests := []struct {
		name     string
		wantErr  bool
		leagueID string
	}{
		{
			name:     "get solo queue",
			leagueID: leagueID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetLeague(tt.leagueID)
			assert.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetStatus(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "get status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetStatus()
			require.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetMatchesByAccount(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
		accountID string
	}{
		{
			name:      "get matches by account",
			accountID: summoner.AccountID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetMatchesByAccount(tt.accountID, 0, 100)
			require.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}

func TestRiotAPIClient_GetMatchesByAccountStream(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
		accountID string
	}{
		{
			name:      "get matches by account stream",
			accountID: summoner.AccountID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.GetMatchesByAccountStream(tt.accountID)
			for res := range got {
				if res.error != nil {
					require.Equal(t, res.error, io.EOF)
					break
				}
				assert.NotNil(t, res.MatchReference)
				fmt.Printf("Match: %#v\n", res.MatchReference)
			}
		})
	}
}

func TestRiotAPIClient_GetMatch(t *testing.T) {
	matches, err := client.GetMatchesByAccount(summoner.AccountID, 0, 1)
	require.Nil(t, err)
	require.Equal(t, len(matches.Matches), 1)
	matchID := matches.Matches[0].GameID
	tests := []struct {
		name    string
		wantErr bool
		matchID int
	}{
		{
			name:    "get match",
			matchID: matchID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetMatch(tt.matchID)
			require.Equal(t, err != nil, tt.wantErr, "error is not nil: ", err)
			if !tt.wantErr {
				assert.NotNil(t, got, "response was nil")
			}
		})
	}
}
