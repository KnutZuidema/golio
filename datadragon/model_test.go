package datadragon

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestChampionData_GetExtended(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		data    *ChampionData
		want    ChampionDataExtended
		wantErr bool
	}
	tests := []test{
		{
			name: "valid",
			doer: mock.NewPathJSONMockDoer(
				[]mock.PathJSONResponse{
					{
						PathSuffix: "/champion.json",
						Object: dataDragonResponse{
							Data: map[string]ChampionData{
								"champion-id": {
									ID:   "champion-id",
									Name: "champion-name",
								},
							},
						},
						Code: 200,
					},
					{
						PathSuffix: "/champion/champion-id.json",
						Object: dataDragonResponse{
							Data: map[string]ChampionDataExtended{
								"champion-id": {
									ChampionData: ChampionData{
										ID:   "champion-id",
										Name: "champion-name",
									},
								},
							},
						},
						Code: 200,
					},
				},
			),
			data: &ChampionData{ID: "champion-id", Name: "champion-name"},
			want: ChampionDataExtended{ChampionData: ChampionData{ID: "champion-id", Name: "champion-name"}},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := test.data.GetExtended(client)
				assert.Equal(t, test.wantErr, err != nil)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestRecommendedItem_GetItem(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		data    *RecommendedItem
		want    Item
		wantErr bool
	}
	tests := []test{
		{
			name: "valid",
			doer: dataDragonResponseDoer(
				map[string]Item{
					"id": {},
				},
			),
			data: &RecommendedItem{ID: "id"},
			want: Item{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionKorea, log.StandardLogger())
				got, err := test.data.GetItem(client)
				assert.Equal(t, test.wantErr, err != nil)
				assert.Equal(t, test.want, got)
			},
		)
	}
}
