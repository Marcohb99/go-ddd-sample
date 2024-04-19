package services

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func initVenuesRecordLabel(t *testing.T) []aggregate.Venue {
	v1, err := aggregate.NewVenue("El Ventorrillo del Cura", "Rincon de la Victoria", 1.1)
	require.NoError(t, err)
	v2, err := aggregate.NewVenue("Sala Velvet", "Málaga", 2.2)
	require.NoError(t, err)
	v3, err := aggregate.NewVenue("Red Rocks Amphitheatre", "Morrison", 3.3)
	require.NoError(t, err)

	return []aggregate.Venue{v1, v2, v3}
}

func TestRecordLabel_BookGigs(t *testing.T) {
	// Given
	venues := initVenuesRecordLabel(t)
	venueIds := []uuid.UUID{venues[0].ID(), venues[1].ID()}

	artist, err := aggregate.NewArtist("Les Claypool", 50)
	require.NoError(t, err)

	gs, err := NewGigService(WithInMemoryVenueRepository(venues), WithInMemoryArtistRepository())
	err = gs.artistRepository.Create(artist)
	require.NoError(t, err)
	sut, err := NewRecordLabel(WithGigService(*gs))
	require.NoError(t, err)

	// When
	err = sut.BookGigs(artist.ID(), venueIds)

	// Then
	require.NoError(t, err)
}
