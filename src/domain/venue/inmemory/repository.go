package inmemory

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/Marcohb99/go-ddd-sample/src/domain/venue"
	"github.com/google/uuid"
	"sync"
)

type VenueRepository struct {
	venues map[uuid.UUID]aggregate.Venue
	sync.Mutex
}

func NewVenueRepository() *VenueRepository {
	return &VenueRepository{
		venues: make(map[uuid.UUID]aggregate.Venue),
	}
}

func (v *VenueRepository) Get(id uuid.UUID) (aggregate.Venue, error) {
	if res, ok := v.venues[id]; ok {
		return res, nil
	}
	return aggregate.Venue{}, venue.ErrVenueNotFound
}
