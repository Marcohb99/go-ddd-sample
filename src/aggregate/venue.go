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
	price    float64
}

func NewVenue(name, location string, price float64) (Venue, error) {
	if name == "" {
		return Venue{}, ErrInvalidVenueName
	}
	return Venue{
		id:       uuid.New(),
		name:     name,
		location: location,
		price:    price,
	}, nil
}

func (v Venue) ID() uuid.UUID {
	return v.id
}

func (v Venue) Price() float64 {
	return v.price
}
