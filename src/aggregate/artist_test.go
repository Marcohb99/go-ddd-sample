package aggregate

import (
	"errors"
	"testing"
)

func Test_NewArtist(t *testing.T) {
	type testCases struct {
		test          string
		name          string
		expectedError error
	}

	tests := []testCases{
		{
			test:          "Empty name validation",
			name:          "",
			expectedError: ErrInvalidPerson,
		},
		{
			test:          "Valid name",
			name:          "John Doe",
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewArtist(tc.name, 30)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
