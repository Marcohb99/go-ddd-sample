package services

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/Marcohb99/go-ddd-sample/src/domain/artist"
	"github.com/Marcohb99/go-ddd-sample/src/domain/artist/inmemory"
	"github.com/Marcohb99/go-ddd-sample/src/domain/venue"
	inmemory2 "github.com/Marcohb99/go-ddd-sample/src/domain/venue/inmemory"
	"github.com/google/uuid"
	"log"
)

type GigConfiguration func(as *GigService) error

// GigService represents the gig service
type GigService struct {
	artistRepository artist.ArtistRepository
	venueRepository  venue.VenueRepository
}

// NewGigService creates a new artist service
func NewGigService(configurations ...GigConfiguration) (*GigService, error) {
	as := &GigService{}

	// loop through the configurations and apply them
	for _, config := range configurations {
		if err := config(as); err != nil {
			return nil, err
		}
	}

	return as, nil
}

func WithArtistRepository(repository artist.ArtistRepository) GigConfiguration {
	return func(as *GigService) error {
		as.artistRepository = repository
		return nil
	}
}

func WithInMemoryArtistRepository() GigConfiguration {
	return func(as *GigService) error {
		repo := inmemory.NewArtistRepository()
		return WithArtistRepository(repo)(as)
	}
}

func WithVenueRepository(repository venue.VenueRepository) GigConfiguration {
	return func(gigService *GigService) error {
		gigService.venueRepository = repository
		return nil
	}
}

func WithInMemoryVenueRepository(venues []aggregate.Venue) GigConfiguration {
	return func(gigService *GigService) error {
		repo := inmemory2.NewVenueRepository()
		for _, v := range venues {
			if err := repo.Add(v); err != nil {
				return err
			}
		}
		return WithVenueRepository(repo)(gigService)
	}
}

func (gs *GigService) BookGigs(artistId uuid.UUID, venueIds []uuid.UUID) (float64, error) {
	bookingArtist, err := gs.artistRepository.Get(artistId)
	if err != nil {
		return 0.0, err
	}

	var venues []aggregate.Venue
	var total float64

	for _, venueId := range venueIds {
		bookingVenue, err := gs.venueRepository.Get(venueId)
		if err != nil {
			return 0.0, err
		}
		venues = append(venues, bookingVenue)
		venueIds = append(venueIds, bookingVenue.ID())
		total += bookingVenue.Price()
	}

	log.Printf("Booking gig for artist %s at venues %v", bookingArtist.Name(), venueIds)
	return total, nil
}
