package services

import (
	"github.com/google/uuid"
	"log"
)

type RecordLabelConfiguration func(rl *RecordLabel) error

type RecordLabel struct {
	gigService GigService
}

func NewRecordLabel(configurations ...RecordLabelConfiguration) (*RecordLabel, error) {
	rl := &RecordLabel{}

	for _, config := range configurations {
		if err := config(rl); err != nil {
			return nil, err
		}
	}

	return rl, nil
}

func WithGigService(gigService GigService) RecordLabelConfiguration {
	return func(rl *RecordLabel) error {
		rl.gigService = gigService
		return nil
	}
}

func (rl *RecordLabel) BookGigs(artistID uuid.UUID, venues []uuid.UUID) error {
	total, err := rl.gigService.BookGigs(artistID, venues)
	if err != nil {
		return err
	}
	log.Printf("Total cost of gigs for artist %s: %f\n", artistID, total)
	return nil
}
