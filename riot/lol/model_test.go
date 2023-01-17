package lol

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/datadragon"
	"github.com/yigithanbalci/golio/internal"
	"github.com/yigithanbalci/golio/internal/mock"
	"github.com/yigithanbalci/golio/static"
)

func TestLeagueList_GetRank(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		entries []*LeagueItem
		want    *LeagueItem
	}{
		{
			name: "get rank",
			i:    0,
			entries: []*LeagueItem{
				{
					LeaguePoints: 0,
				},
				{
					LeaguePoints: 1,
				},
			},
			want: &LeagueItem{
				LeaguePoints: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LeagueList{
				Entries: tt.entries,
			}
			require.Equal(t, tt.want, l.GetRank(tt.i))
		})
	}
}

func TestChampionInfo_GetChampionsForNewPlayers(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionInfo
		want    []datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion1": {Key: "1", Name: "champion1"},
				"champion2": {Key: "2", Name: "champion2"},
			}),
			model: ChampionInfo{
				FreeChampionIDsForNewPlayers: []int{1, 2},
			},
			want: []datadragon.ChampionDataExtended{
				{ChampionData: datadragon.ChampionData{Name: "champion1", Key: "1"}},
				{ChampionData: datadragon.ChampionData{Name: "champion2", Key: "2"}},
			},
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			model: ChampionInfo{
				FreeChampionIDsForNewPlayers: []int{1},
			},
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampionsForNewPlayers(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestChampionInfo_GetChampions(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionInfo
		want    []datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion1": {Key: "1", Name: "champion1"},
				"champion2": {Key: "2", Name: "champion2"},
			}),
			model: ChampionInfo{
				FreeChampionIDs: []int{1, 2},
			},
			want: []datadragon.ChampionDataExtended{
				{ChampionData: datadragon.ChampionData{Name: "champion1", Key: "1"}},
				{ChampionData: datadragon.ChampionData{Name: "champion2", Key: "2"}},
			},
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			model: ChampionInfo{
				FreeChampionIDs: []int{1},
			},
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampions(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestChampionMastery_GetSummoner(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionMastery
		want    *Summoner
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Summoner{ID: "id"}, 200),
			model: ChampionMastery{SummonerID: "id"},
			want:  &Summoner{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
			got, err := test.model.GetSummoner(NewClient(client))
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestChampionMastery_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   ChampionMastery
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion": {Name: "champion", Key: "1"},
			}),
			model: ChampionMastery{ChampionID: 1},
			want:  datadragon.ChampionDataExtended{ChampionData: datadragon.ChampionData{Name: "champion", Key: "1"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampion(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestLeagueItem_GetSummoner(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   LeagueItem
		want    *Summoner
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Summoner{ID: "id"}, 200),
			model: LeagueItem{SummonerID: "id"},
			want:  &Summoner{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
			got, err := test.model.GetSummoner(NewClient(client))
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMatchInfo_GetQueue(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.Queue
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.Queue{{ID: 1}}, 200),
			model: MatchInfo{QueueID: 1},
			want:  static.Queue{ID: 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := static.NewClient(test.doer, log.StandardLogger())
			got, err := test.model.GetQueue(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMatchInfo_GetMap(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.Map
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.Map{{ID: 1}}, 200),
			model: MatchInfo{MapID: 1},
			want:  static.Map{ID: 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := static.NewClient(test.doer, log.StandardLogger())
			got, err := test.model.GetMap(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMatchInfo_GetGameType(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.GameType
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.GameType{{Type: "type"}}, 200),
			model: MatchInfo{GameType: "type"},
			want:  static.GameType{Type: "type"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := static.NewClient(test.doer, log.StandardLogger())
			got, err := test.model.GetGameType(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMatchInfo_GetGameMode(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchInfo
		want    static.GameMode
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer([]static.GameMode{{Mode: "type"}}, 200),
			model: MatchInfo{GameMode: "type"},
			want:  static.GameMode{Mode: "type"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := static.NewClient(test.doer, log.StandardLogger())
			got, err := test.model.GetGameMode(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetSummoner(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    *Summoner
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Summoner{ID: "id"}, 200),
			model: Participant{SummonerID: "id"},
			want:  &Summoner{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
			got, err := test.model.GetSummoner(NewClient(client))
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetProfileIcon(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.ProfileIcon
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ProfileIcon{
				"champion": {ID: 1},
			}),
			model: Participant{ProfileIcon: 1},
			want:  datadragon.ProfileIcon{ID: 1},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetProfileIcon(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestTeamBan_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   TeamBan
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion": {Name: "champion", Key: "1"},
			}),
			model: TeamBan{ChampionID: 1},
			want:  datadragon.ChampionDataExtended{ChampionData: datadragon.ChampionData{Name: "champion", Key: "1"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampion(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion": {Name: "champion", Key: "1"},
			}),
			model: Participant{ChampionID: 1},
			want:  datadragon.ChampionDataExtended{ChampionData: datadragon.ChampionData{Name: "champion", Key: "1"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampion(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetSpell1(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.SummonerSpell{
				"champion": {Key: "1"},
			}),
			model: Participant{Summoner1ID: 1},
			want:  datadragon.SummonerSpell{Key: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetSpell1(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetSpell2(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.SummonerSpell{
				"champion": {Key: "2"},
			}),
			model: Participant{Summoner2ID: 2},
			want:  datadragon.SummonerSpell{Key: "2"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetSpell2(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem0(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item0: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem0(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem1(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item1: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem1(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem2(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item2: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem2(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem3(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item3: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem3(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem4(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item4: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem4(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem5(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item5: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem5(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestParticipant_GetItem6(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   Participant
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: Participant{Item6: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem6(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestBannedChampion_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   BannedChampion
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion": {Name: "champion", Key: "1"},
			}),
			model: BannedChampion{ChampionID: 1},
			want:  datadragon.ChampionDataExtended{ChampionData: datadragon.ChampionData{Name: "champion", Key: "1"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampion(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCurrentGameParticipant_GetChampion(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   CurrentGameParticipant
		want    datadragon.ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.ChampionData{
				"champion": {Name: "champion", Key: "1"},
			}),
			model: CurrentGameParticipant{ChampionID: 1},
			want:  datadragon.ChampionDataExtended{ChampionData: datadragon.ChampionData{Name: "champion", Key: "1"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetChampion(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCurrentGameParticipant_GetSpell1(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   CurrentGameParticipant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.SummonerSpell{
				"champion": {Key: "1"},
			}),
			model: CurrentGameParticipant{Spell1ID: 1},
			want:  datadragon.SummonerSpell{Key: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetSpell1(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCurrentGameParticipant_GetSpell2(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   CurrentGameParticipant
		want    datadragon.SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.SummonerSpell{
				"champion": {Key: "2"},
			}),
			model: CurrentGameParticipant{Spell2ID: 2},
			want:  datadragon.SummonerSpell{Key: "2"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetSpell2(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGameInfo_GetMatch(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   GameInfo
		want    *Match
		wantErr error
	}
	tests := []test{
		{
			name:  "valid",
			doer:  mock.NewJSONMockDoer(Match{Metadata: &MatchMetadata{MatchID: "NA1_1"}}, 200),
			model: GameInfo{GameID: 1},
			want:  &Match{Metadata: &MatchMetadata{MatchID: "NA1_1"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionKorea, "key", test.doer, log.StandardLogger())
			got, err := test.model.GetMatch(NewClient(client))
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMatchEvent_GetItem(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		model   MatchEvent
		want    datadragon.Item
		wantErr error
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(map[string]datadragon.Item{
				"1": {},
			}),
			model: MatchEvent{ItemID: 1},
			want:  datadragon.Item{ID: "1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := datadragon.NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.model.GetItem(client)
			assert.Equal(t, test.wantErr, err)
			assert.Equal(t, test.want, got)
		})
	}
}

type dataDragonResponse struct {
	Type    string
	Format  string
	Version string
	Data    interface{}
}

func dataDragonResponseDoer(object interface{}) internal.Doer {
	return mock.NewJSONMockDoer(dataDragonResponse{
		Data: object,
	}, 200)
}
