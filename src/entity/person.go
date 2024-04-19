package entity

import "github.com/google/uuid"

// Person represents a person entity
type Person struct {
	ID   uuid.UUID
	Name string
	Age  int
}
