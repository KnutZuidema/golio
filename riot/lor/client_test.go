package lor

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
	"github.com/yigithanbalci/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	c := NewClient(internal.NewClient(api.RegionBrasil, "key", mock.NewStatusMockDoer(200), logrus.StandardLogger()))
	if c == nil {
		t.Error("returned nil")
	}
}
