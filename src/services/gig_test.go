package services

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func init_venues(t *testing.T) []aggregate.Venue {
	v1, err := aggregate.NewVenue("El Ventorrillo del Cura", "Rincon de la Victoria", 1.1)
	require.NoError(t, err)
	v2, err := aggregate.NewVenue("Sala Velvet", "Málaga", 2.2)
	require.NoError(t, err)
	v3, err := aggregate.NewVenue("Red Rocks Amphitheatre", "Morrison", 3.3)
	require.NoError(t, err)

	return []aggregate.Venue{v1, v2, v3}
}

func Test_NewGigs(t *testing.T) {
	venues := init_venues(t)
	sut, err := NewGigService(WithInMemoryArtistRepository(), WithInMemoryVenueRepository(venues))
	require.NoError(t, err)

	artist, err := aggregate.NewArtist("Les Claypool", 50)
	err = sut.artistRepository.Create(artist)
	require.NoError(t, err)

	productIds := []uuid.UUID{venues[0].ID(), venues[1].ID()}
	total, err := sut.BookGigs(artist.ID(), productIds)
	require.NoError(t, err)
	assert.Equal(t, 3.3000000000000003, total)
}
