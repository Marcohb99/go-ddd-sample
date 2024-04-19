package valueobject

import "github.com/google/uuid"

// Collab represents a collaboration value object. It is a VO because it cannot be transformed
type Collab struct {
	mainArtist uuid.UUID
	featArtist uuid.UUID
	date       string
}
