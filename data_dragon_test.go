package golio

import (
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var (
	ddClient *DataDragonClient
)

func TestNewDataDragonClient(t *testing.T) {
	ddClient = NewDataDragonClient(http.DefaultClient, RegionEuropeWest, log.StandardLogger())
	require.NotNil(t, ddClient)
}

func TestDataDragonClient_GetChampions(t *testing.T) {
	got, err := ddClient.GetChampions()
	require.Nil(t, err)
	require.NotNil(t, got)
}

func TestDataDragonClient_GetChampion(t *testing.T) {
	got, err := ddClient.GetChampion("Ashe")
	require.Nil(t, err)
	require.NotNil(t, got)
}

func TestDataDragonClient_GetProfileIcons(t *testing.T) {
	got, err := ddClient.GetProfileIcons()
	require.Nil(t, err)
	require.NotNil(t, got)
}

func TestDataDragonClient_GetItems(t *testing.T) {
	got, err := ddClient.GetItems()
	require.Nil(t, err)
	require.NotNil(t, got)
}

func TestDataDragonClient_GetMasteries(t *testing.T) {
	got, err := ddClient.GetMasteries()
	require.Nil(t, err)
	require.NotNil(t, got)
}

func TestDataDragonClient_GetSummonerSpells(t *testing.T) {
	got, err := ddClient.GetSummonerSpells()
	require.Nil(t, err)
	require.NotNil(t, got)
}
