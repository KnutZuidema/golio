package lol

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"strconv"

	"github.com/KnutZuidema/golio/internal"
)

// MatchClient provides methods for the match endpoints of the League of Legends API.
type MatchClient struct {
	c *internal.Client
}

// MatchListOptions
type MatchListOptions struct {
	// Set of champion IDs for filtering the matchlist.
	Champion int
	// Set of queue IDs for filtering the matchlist.
	Queue int
	// The begin time to use for filtering matchlist specified as epoch milliseconds.
	BeginTime int64
	// The end time to use for filtering matchlist specified as epoch milliseconds.
	EndTime int64
}

func (mo *MatchListOptions) buildParam() string {
	var param string
	if mo.Champion != 0 {
		param += "&champion=" + strconv.Itoa(mo.Champion)
	}
	if mo.Queue != 0 {
		param += "&queue=" + strconv.Itoa(mo.Queue)
	}
	if mo.BeginTime != 0 {
		param += "&beginTime=" + strconv.FormatInt(mo.BeginTime, 10)
	}
	if mo.EndTime != 0 {
		param += "&endTime=" + strconv.FormatInt(mo.EndTime, 10)
	}
	return param
}

// Get returns a match specified by its ID
func (m *MatchClient) Get(id int) (*Match, error) {
	logger := m.logger().WithField("method", "Get")
	var match *Match
	if err := m.c.GetInto(fmt.Sprintf(endpointGetMatch, id), &match); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return match, nil
}

// List returns a specified range of matches played on the account
func (m *MatchClient) List(accountID string, beginIndex, endIndex int, options ...*MatchListOptions) (*Matchlist, error) {
	logger := m.logger().WithField("method", "List")
	var matches *Matchlist
	endpoint := fmt.Sprintf(endpointGetMatchesByAccount, accountID, beginIndex, endIndex)
	if len(options) != 0 {
		endpoint += options[0].buildParam()
	}
	if err := m.c.GetInto(endpoint, &matches); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return matches, nil
}

// MatchStreamValue value returned by ListStream, containing either a reference to a match or an error
type MatchStreamValue struct {
	*MatchReference
	Error error
}

// ListStream returns all matches played on this account as a stream, requesting new until there are no
// more new games
func (m *MatchClient) ListStream(accountID string) <-chan MatchStreamValue {
	logger := m.logger().WithField("method", "ListStream")
	cMatches := make(chan MatchStreamValue, 100)
	go func() {
		start := 0
		for {
			matches, err := m.List(accountID, start, start+100)
			if err != nil {
				logger.Debug(err)
				cMatches <- MatchStreamValue{Error: err}
				return
			}
			for _, match := range matches.Matches {
				cMatches <- MatchStreamValue{MatchReference: match}
			}
			if len(matches.Matches) < 100 {
				cMatches <- MatchStreamValue{Error: io.EOF}
				return
			}
			start += 100
		}
	}()
	return cMatches
}

// GetTimeline returns the timeline for the given match
// NOTE: timelines are not available for every match
func (m *MatchClient) GetTimeline(matchID int) (*MatchTimeline, error) {
	logger := m.logger().WithField("method", "GetTimeline")
	var timeline MatchTimeline
	if err := m.c.GetInto(fmt.Sprintf(endpointGetMatchTimeline, matchID), &timeline); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &timeline, nil
}

// ListIDsByTournamentCode returns all match ids for the given tournament
func (m *MatchClient) ListIDsByTournamentCode(tournamentCode string) ([]int, error) {
	logger := m.logger().WithField("method", "ListIDsByTournamentCode")
	var ids []int
	if err := m.c.GetInto(fmt.Sprintf(endpointGetMatchIDsByTournamentCode, tournamentCode), &ids); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return ids, nil
}

// GetForTournament returns the match data for the given match in the given tournament
func (m *MatchClient) GetForTournament(matchID int, tournamentCode string) (*Match, error) {
	logger := m.logger().WithField("method", "GetForTournament")
	var match Match
	if err := m.c.GetInto(
		fmt.Sprintf(endpointGetMatchForTournament, matchID, tournamentCode),
		&match,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &match, nil
}

func (m *MatchClient) logger() log.FieldLogger {
	return m.c.Logger().WithField("category", "match")
}
