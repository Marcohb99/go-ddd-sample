package artist

import (
	"errors"
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
)

var (
	ErrArtistNotFound = errors.New("artist not found")
	ErrFailedToCreate = errors.New("failed to create artist")
	ErrFailedToUpdate = errors.New("failed to update artist")
)

type ArtistRepository interface {
	Get(uuid.UUID) (aggregate.Artist, error)
	Create(aggregate.Artist) error
	Update(aggregate.Artist) error
}
