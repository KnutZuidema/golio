package golio

import (
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
)

func TestNewClient(t *testing.T) {
	client := NewClient(
		"api_key",
		WithLogger(log.StandardLogger()),
		WithRegion(api.RegionEuropeWest),
		WithClient(http.DefaultClient),
	)
	require.NotNil(t, client)
}
