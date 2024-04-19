package venue

import (
	"errors"
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
)

var (
	ErrVenueNotFound = errors.New("venue not found")
	ErrVenueExists   = errors.New("venue already exists")
)

type VenueRepository interface {
	Get(uuid.UUID) (aggregate.Venue, error)
	Add(aggregate.Venue) error
}
