package inmemory

import (
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/google/uuid"
	"sync"
)

type ArtistRepository struct {
	artists map[uuid.UUID]aggregate.Artist
	sync.Mutex
}

func NewArtistRepository() *ArtistRepository {
	return &ArtistRepository{
		artists: make(map[uuid.UUID]aggregate.Artist),
	}
}

func (r *ArtistRepository) Get(id uuid.UUID) (aggregate.Artist, error) {
	return aggregate.Artist{}, nil
}

func (r *ArtistRepository) Create(a aggregate.Artist) error {
	return nil
}

func (r *ArtistRepository) Update(a aggregate.Artist) error {
	return nil
}
