package venue

import (
	"errors"
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
)

var (
	ErrVenueNotFound = errors.New("venue not found")
)

type VenueRepository interface {
	Get(uuid.UUID) (aggregate.Venue, error)
}
