package inmemory

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	venueDomain "github.com/Marcohb99/go-ddd-sample/src/domain/venue"
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
	return aggregate.Venue{}, venueDomain.ErrVenueNotFound
}

func (v *VenueRepository) Add(venue aggregate.Venue) error {
	if v.venues == nil {
		v.Lock()
		v.venues = make(map[uuid.UUID]aggregate.Venue)
		v.Unlock()
	}
	v.Lock()
	defer v.Unlock()
	if _, ok := v.venues[venue.ID()]; ok {
		return venueDomain.ErrVenueExists
	}
	v.venues[venue.ID()] = venue
	return nil
}
