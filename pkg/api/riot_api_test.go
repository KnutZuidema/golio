package api

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/KnutZuidema/golio/pkg/model"
)

const (
	apiKey     = "RGAPI-70e73c6f-8ff5-4bee-8029-32968bdfd922"
	testRegion = RegionEuropeWest
)

var (
	client               *Client
	summonerByRegion     = map[region]*model.Summoner{}
	summonerNameByRegion = map[region]string{
		RegionEuropeWest:        "SK Jenax",
		RegionEuropeNorthEast:   "I am LeBron",
		RegionTurkey:            "Reformed Hatred",
		RegionRussia:            "The Great Donald",
		RegionOceania:           "k1ngggggggg",
		RegionNorthAmerica:      "tarzaned5",
		RegionLatinAmericaSouth: "CodyStark",
		RegionLatinAmericaNorth: "Shym",
		RegionJapan:             "isurugi",
		RegionBrasil:            "paiN 25789",
		RegionKorea:             "Cuzz",
	}
)

func init() {
	client = NewClient(RegionEuropeWest, apiKey, http.DefaultClient, logrus.StandardLogger())
	for reg, summoner := range summonerNameByRegion {
		client.Region = reg
		s, err := client.GetSummonerByName(summoner)
		if err != nil {
			panic(err)
		}
		summonerByRegion[reg] = s
	}
	client.Region = testRegion
}

func TestRiotAPIClient_GetSummonerByName(t *testing.T) {
	tests := []struct {
		name         string
		summonerName string
		wantErr      bool
		want         *model.Summoner
	}{
		{
			name:         "get summoner",
			summonerName: summonerNameByRegion[testRegion],
			want:         summonerByRegion[testRegion],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			got, err := client.GetSummonerByName(tt.summonerName)
			assert.Equal(t, err != nil, tt.wantErr, fmt.Sprintf("error is not %v: %v", tt.wantErr, err))
			if !tt.wantErr {
				assert.Equal(t, got, tt.want, "responses does not match")
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
			accountID: summonerByRegion[testRegion].AccountID,
			want:      summonerByRegion[testRegion],
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
			puuid: summonerByRegion[testRegion].PUUID,
			want:  summonerByRegion[testRegion],
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
			summonerID: summonerByRegion[testRegion].ID,
			want:       summonerByRegion[testRegion],
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
			summonerID: summonerByRegion[testRegion].ID,
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
	champion, err := client.GetChampion("Ashe")
	require.Nil(t, err)
	tests := []struct {
		name       string
		summonerID string
		championID string
		wantErr    bool
	}{
		{
			name:       "get champion mastery",
			summonerID: summonerByRegion[testRegion].ID,
			championID: champion.Key,
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
			summonerID: summonerByRegion[testRegion].ID,
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
			summonerID: summonerByRegion[testRegion].ID,
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
	tests := []struct {
		name     string
		wantErr  bool
		queue    queue
		tier     tier
		division division
	}{
		{
			name:     "get leagues",
			queue:    QueueRankedSolo,
			tier:     TierGold,
			division: DivisionThree,
		},
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
			accountID: summonerByRegion[testRegion].AccountID,
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
			accountID: summonerByRegion[testRegion].AccountID,
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
			}
		})
	}
}

func TestRiotAPIClient_GetMatch(t *testing.T) {
	matches, err := client.GetMatchesByAccount(summonerByRegion[testRegion].AccountID, 0, 1)
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
