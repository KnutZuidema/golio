package datadragon

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
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
			doer: dataDragonResponseDoer(map[string]*ChampionDataExtended{
				"champion": {ChampionData: ChampionData{Name: "champion"}},
			}),
			data: &ChampionData{Name: "champion"},
			want: ChampionDataExtended{ChampionData: ChampionData{Name: "champion"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
			got, err := test.data.GetExtended(client)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want, got)
		})
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
			doer: dataDragonResponseDoer(map[string]Item{
				"id": {},
			}),
			data: &RecommendedItem{ID: "id"},
			want: Item{ID: "id"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := NewClient(test.doer, api.RegionKorea, log.StandardLogger())
			got, err := test.data.GetItem(client)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.want, got)
		})
	}
}
