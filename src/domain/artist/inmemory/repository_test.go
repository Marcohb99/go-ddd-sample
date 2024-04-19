package inmemory

import (
	"errors"
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	artistDomain "github.com/Marcohb99/go-ddd-sample/src/domain/artist"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InMemoryRepository_GetCustomer(t *testing.T) {
	type testCase struct {
		test          string
		id            uuid.UUID
		expectedError error
	}

	artist, err := aggregate.NewArtist("Les Claypool", 50)
	require.NoError(t, err)

	sut := ArtistRepository{
		artists: map[uuid.UUID]aggregate.Artist{
			artist.ID(): artist,
		},
	}

	tests := []testCase{
		{
			test:          "Get non-existing customer",
			id:            uuid.New(),
			expectedError: artistDomain.ErrArtistNotFound,
		},
		{
			test:          "Get existing customer",
			id:            artist.ID(),
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			_, err := sut.Get(tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
