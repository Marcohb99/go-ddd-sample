package aggregate

import (
	"errors"
	"testing"
)

func Test_NewVenue(t *testing.T) {
	type testCases struct {
		test          string
		name          string
		location      string
		expectedError error
	}
	tests := []testCases{
		{
			test:          "Empty name validation",
			name:          "",
			location:      "Rincon de la Victoria",
			expectedError: ErrInvalidVenueName,
		},
		{
			test:          "Valid name",
			name:          "El Ventorrillo del Cura",
			location:      "Rincon de la Victoria",
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewVenue(tc.name, tc.location, 1.1)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
