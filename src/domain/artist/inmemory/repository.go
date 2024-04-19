package inmemory

import (
	"fmt"
	"github.com/Marcohb99/go-ddd-sample/src/aggregate"
	"github.com/Marcohb99/go-ddd-sample/src/domain/artist"
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
	r.Lock()
	defer r.Unlock()

	if result, ok := r.artists[id]; ok {
		return result, nil
	}

	return aggregate.Artist{}, artist.ErrArtistNotFound
}

func (r *ArtistRepository) Create(a aggregate.Artist) error {
	r.ensureArtistMapInitialized()
	if _, ok := r.artists[a.ID()]; ok {
		return fmt.Errorf("artist with ID %s already exists", artist.ErrFailedToCreate)
	}
	r.Lock()
	r.artists[a.ID()] = a
	r.Unlock()
	return nil
}

func (r *ArtistRepository) ensureArtistMapInitialized() {
	if r.artists == nil {
		r.Lock()
		r.artists = make(map[uuid.UUID]aggregate.Artist)
		r.Unlock()
	}
}

func (r *ArtistRepository) Update(a aggregate.Artist) error {
	r.ensureArtistMapInitialized()
	if _, ok := r.artists[a.ID()]; !ok {
		return fmt.Errorf("artist with ID %s not found", artist.ErrFailedToUpdate)
	}
	r.Lock()
	r.artists[a.ID()] = a
	r.Unlock()
	return nil
}
