package services

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func init_venues(t *testing.T) []aggregate.Venue {
	v1, err := aggregate.NewVenue("El Ventorrillo del Cura", "Rincon de la Victoria")
	require.NoError(t, err)
	v2, err := aggregate.NewVenue("Sala Velvet", "Málaga")
	require.NoError(t, err)
	v3, err := aggregate.NewVenue("Red Rocks Amphitheatre", "Morrison")
	require.NoError(t, err)

	return []aggregate.Venue{v1, v2, v3}
}

func Test_NewGig(t *testing.T) {
	venues := init_venues(t)
	sut, err := NewGigService(WithInMemoryArtistRepository(), WithInMemoryVenueRepository(venues))
	require.NoError(t, err)

	artist, err := aggregate.NewArtist("Les Claypool", 50)
	err = sut.artistRepository.Create(artist)
	require.NoError(t, err)

	productIds := []uuid.UUID{venues[0].ID(), venues[1].ID()}
	err = sut.BookGig(artist.ID(), productIds)
	require.NoError(t, err)
}
