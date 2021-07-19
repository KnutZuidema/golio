package lol

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
)

// MatchClient provides methods for the match endpoints of the League of Legends API.
type MatchClient struct {
	c *internal.Client
}

// MatchListOptions providing additional options for List
type MatchListOptions struct {
	// The queue ID for filtering the matchlist.
	Queue *int
	// The game type for filtering the matchlist (see static.GameType.Type).
	Type string
}

func (mo *MatchListOptions) buildParam() string {
	var param string
	if mo.Queue != nil {
		param += "&queue=" + strconv.Itoa(*mo.Queue)
	}
	if mo.Type != "" {
		param += "&type=" + mo.Type
	}
	return param
}

// Get returns a match specified by its ID
func (m *MatchClient) Get(id string) (*Match, error) {
	logger := m.logger().WithField("method", "Get")
	c := *m.c                                          // copy client
	c.Region = api.Region(api.RegionToRoute[c.Region]) // Match v5 uses a route instead of a region
	var match *Match
	if err := c.GetInto(fmt.Sprintf(endpointGetMatch, id), &match); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return match, nil
}

// List returns a specified list of match IDs played on the account
func (m *MatchClient) List(puuid string, start, count int, options ...*MatchListOptions) (
	[]string, error) {
	logger := m.logger().WithField("method", "List")
	c := *m.c                                          // copy client
	c.Region = api.Region(api.RegionToRoute[c.Region]) // Match v5 uses a route instead of a region
	var matches []string
	endpoint := fmt.Sprintf(endpointGetMatchIDs, puuid, start, count)
	if len(options) != 0 {
		endpoint += options[0].buildParam()
	}
	if err := c.GetInto(endpoint, &matches); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return matches, nil
}

// MatchStreamValue value returned by ListStream, containing either a reference to a match or an error
type MatchStreamValue struct {
	MatchID string
	Error   error
}

// ListStream returns all matches played on this account as a stream, requesting new until there are no
// more new games
func (m *MatchClient) ListStream(puuid string, options ...*MatchListOptions) <-chan MatchStreamValue {
	logger := m.logger().WithField("method", "ListStream")
	cMatches := make(chan MatchStreamValue, 100)

	// Copy the input options to prevent caller modification while streaming
	opts := make([]*MatchListOptions, 0)
	for _, o := range options {
		// Copy the value in case the caller modifies it after we return
		queue := *options[0].Queue
		// Shallow copy the other values
		newOpt := *o
		newOpt.Queue = &queue
		opts = append(opts, &newOpt)
	}
	if len(options) != 0 && options[0].Queue != nil {
		// Copy the value in case the caller modifies it after we return
		queue := *options[0].Queue
		options[0].Queue = &queue
	}
	go func() {
		defer close(cMatches)
		start := 0
		for {
			matches, err := m.List(puuid, start, 100, opts...)
			if err != nil {
				logger.Debug(err)
				cMatches <- MatchStreamValue{Error: err}
				return
			}
			for _, match := range matches {
				cMatches <- MatchStreamValue{MatchID: match}
			}
			if len(matches) < 100 {
				return
			}
			start += 100
		}
	}()
	return cMatches
}

// GetTimeline returns the timeline for the given match
// NOTE: timelines are not available for every match
// TODO: update to v5 when struct is documented
func (m *MatchClient) GetTimeline(matchID int) (*MatchTimeline, error) {
	logger := m.logger().WithField("method", "GetTimeline")
	var timeline MatchTimeline
	if err := m.c.GetInto(fmt.Sprintf(endpointGetMatchTimeline, fmt.Sprint(matchID)), &timeline); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &timeline, nil
}

func (m *MatchClient) logger() log.FieldLogger {
	return m.c.Logger().WithField("category", "match")
}
