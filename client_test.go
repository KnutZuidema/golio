package golio

import (
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewClient(RegionEuropeWest, os.Getenv("API_KEY"), http.DefaultClient, logrus.StandardLogger())
	require.NotNil(t, client)
}
