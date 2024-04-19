package inmemory

import (
	"errors"
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	venueRepository "github.com/Marcohb99/go-ddd-sample/src/domain/venue"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Repository_Get_Venue(t *testing.T) {
	venue, err := aggregate.NewVenue("El Ventorrillo del Cura", "Rincon de la Victoria")
	require.NoError(t, err)

	sut := VenueRepository{
		venues: map[uuid.UUID]aggregate.Venue{
			venue.ID(): venue,
		},
	}

	type testCase struct {
		name          string
		id            uuid.UUID
		expectedVenue aggregate.Venue
		expectedError error
	}

	tests := []testCase{
		{
			name:          "Given a valid venue ID, it returns the venue",
			id:            venue.ID(),
			expectedVenue: venue,
			expectedError: nil,
		},
		{
			name:          "Given an invalid venue ID, it returns an error",
			id:            uuid.New(),
			expectedVenue: aggregate.Venue{},
			expectedError: venueRepository.ErrVenueNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			venue, err := sut.Get(tt.id)
			assert.Equal(t, tt.expectedVenue, venue)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

func Test_Repository_Add_Venue(t *testing.T) {
	venue, err := aggregate.NewVenue("El Ventorrillo del Cura", "Rincon de la Victoria")
	require.NoError(t, err)

	sut := VenueRepository{
		venues: map[uuid.UUID]aggregate.Venue{
			venue.ID(): venue,
		},
	}

	type testCase struct {
		name          string
		venue         aggregate.Venue
		expectedError error
		expectedCount int
	}

	v, err := aggregate.NewVenue("Sala Velvet", "Málaga")
	require.NoError(t, err)

	// Note, if we switch the order, the first one will add the venue
	// and the second one will fail because the expected count is 1
	tests := []testCase{
		{
			name:          "Given an invalid venue ID, it returns an error",
			venue:         venue,
			expectedError: venueRepository.ErrVenueExists,
			expectedCount: 1,
		},
		{
			name:          "Given a valid venue, it is correctly added",
			venue:         v,
			expectedError: nil,
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := sut.Add(tt.venue)
			assert.Equal(t, tt.expectedCount, len(sut.venues))
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
