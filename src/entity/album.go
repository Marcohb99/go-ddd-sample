package entity

import (
	"github.com/google/uuid"
)

// Album represents an album entity
type Album struct {
	ID          uuid.UUID
	Name        string
	ReleaseDate string
}
