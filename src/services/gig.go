package services

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/Marcohb99/go-ddd-sample/src/domain/artist"
	"github.com/Marcohb99/go-ddd-sample/src/domain/artist/inmemory"
	"github.com/Marcohb99/go-ddd-sample/src/domain/venue"
	"github.com/google/uuid"
	"log"
)

type GigConfiguration func(as *GigService) error

// GigService represents the artist service
type GigService struct {
	artistRepository artist.ArtistRepository
	venueRepository  venue.VenueRepository
}

// NewArtistService creates a new artist service
func NewArtistService(configurations ...GigConfiguration) (*GigService, error) {
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

func (gs *GigService) BookGig(artistId uuid.UUID, venueIds []uuid.UUID) error {
	bookingArtist, err := gs.artistRepository.Get(artistId)
	if err != nil {
		return err
	}
	var venues []aggregate.Venue
	for _, venueId := range venueIds {
		bookingVenue, err := gs.venueRepository.Get(venueId)
		if err != nil {
			return err
		}
		venues = append(venues, bookingVenue)
	}
	log.Printf("Booking gig for artist %s at venues %v", bookingArtist.Name(), venues)
	return nil
}