package aggregate

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidVenueName = errors.New("invalid venue name")
)

type Venue struct {
	id       uuid.UUID
	name     string
	location string
}

func NewVenue(name, location string) (Venue, error) {
	if name == "" {
		return Venue{}, ErrInvalidVenueName
	}
	return Venue{
		id:       uuid.New(),
		name:     name,
		location: location,
	}, nil
}

func (v Venue) ID() uuid.UUID {
	return v.id
}
