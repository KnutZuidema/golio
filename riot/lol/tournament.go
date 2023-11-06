package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// TournamentClient provides methods for the tournament endpoints of the League of Legends API.
type TournamentClient struct {
	c *internal.Client
}

// CreateCodes creates a specified amount of codes for a tournament.
// For more information about the parameters see the documentation for TournamentCodeParameters.
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (t *TournamentClient) CreateCodes(id, count int, params *TournamentCodeParameters, stub bool) ([]string, error) {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "CreateCodes",
			"stub":   stub,
		},
	)
	endpoint := endpointCreateTournamentCodes
	if stub {
		endpoint = endpointCreateStubTournamentCodes
	}
	var codes []string
	if err := t.c.PostInto(fmt.Sprintf(endpoint, count, id), params, &codes); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return codes, nil
}

// ListLobbyEvents returns the lobby events for a lobby specified by the tournament code
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (t *TournamentClient) ListLobbyEvents(code string, useStub bool) (*LobbyEventList, error) {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "ListLobbyEvents",
			"stub":   useStub,
		},
	)
	endpoint := endpointGetLobbyEvents
	if useStub {
		endpoint = endpointGetStubLobbyEvents
	}
	var events LobbyEventList
	if err := t.c.GetInto(fmt.Sprintf(endpoint, code), &events); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &events, nil
}

// CreateProvider creates a tournament provider and returns the ID.
// For more information about the parameters see the documentation for ProviderRegistrationParameters.
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (t *TournamentClient) CreateProvider(parameters *ProviderRegistrationParameters, useStub bool) (int, error) {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "CreateProvider",
			"stub":   useStub,
		},
	)
	endpoint := endpointCreateTournamentProvider
	if useStub {
		endpoint = endpointCreateStubTournamentProvider
	}
	var id int
	if err := t.c.PostInto(endpoint, parameters, &id); err != nil {
		logger.Debug(err)
		return 0, err
	}
	return id, nil
}

// Create creates a tournament and returns the ID.
// For more information about the parameters see the documentation for TournamentRegistrationParameters.
// Set the useStub flag to true to use the stub endpoints for mocking an implementation
func (t *TournamentClient) Create(parameters *TournamentRegistrationParameters, useStub bool) (int, error) {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "Create",
			"stub":   useStub,
		},
	)
	endpoint := endpointCreateTournament
	if useStub {
		endpoint = endpointCreateStubTournament
	}
	var id int
	if err := t.c.PostInto(endpoint, parameters, &id); err != nil {
		logger.Debug(err)
		return 0, err
	}
	return id, nil
}

// Get returns an existing tournament
func (t *TournamentClient) Get(code string) (*Tournament, error) {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "Get",
		},
	)
	var tournament Tournament
	if err := t.c.GetInto(fmt.Sprintf(endpointGetTournament, code), &tournament); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &tournament, nil
}

// Update updates an existing tournament
func (t *TournamentClient) Update(code string, parameters TournamentUpdateParameters) error {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "Update",
		},
	)
	if err := t.c.Put(fmt.Sprintf(endpointUpdateTournament, code), parameters); err != nil {
		logger.Debug(err)
		return err
	}
	return nil
}

func (t *TournamentClient) logger() log.FieldLogger {
	return t.c.Logger().WithField("category", "tournament")
}
