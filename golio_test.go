package golio

import (
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
)

func TestNewClient(t *testing.T) {
	client := NewClient(api.RegionEuropeWest, os.Getenv("API_KEY"), http.DefaultClient, logrus.StandardLogger())
	require.NotNil(t, client)
}
