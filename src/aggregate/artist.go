package aggregate

import (
	"errors"
	"github.com/Marcohb99/go-ddd-sample/src/entity"
	"github.com/Marcohb99/go-ddd-sample/src/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("invalid name")
)

// Artist represents an artist aggregate
type Artist struct {
	person  *entity.Person
	albums  []*entity.Album
	collabs []*valueobject.Collab
}

func NewArtist(name string, age int) (Artist, error) {
	if name == "" {
		return Artist{}, ErrInvalidPerson
	}
	return Artist{
		person: &entity.Person{
			ID:   uuid.New(),
			Name: name,
			Age:  age,
		},
		albums:  make([]*entity.Album, 0),
		collabs: make([]*valueobject.Collab, 0),
	}, nil
}

func (a *Artist) ID() uuid.UUID {
	return a.person.ID
}

func (a *Artist) Name() string {
	return a.person.Name
}

func (a *Artist) Age() int {
	return a.person.Age
}

func (a *Artist) Albums() []*entity.Album {
	return a.albums
}

func (a *Artist) Collabs() []*valueobject.Collab {
	return a.collabs
}
