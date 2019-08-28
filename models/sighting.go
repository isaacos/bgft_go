package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
)

// Sighting model struct
type Sighting struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Location	string 		`json:"location" db:"location"`
	Description	string 	`json:"description" db:"description"`
	Lat 				float64	`json:"lat" db:"lat"`
	Lng 				float64	`json:"lng" db:"lng"`
	User				*User 	`belongs_to:"user" json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (s Sighting) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Sightings is not required by pop and may be deleted
type Sightings []Sighting

// String is not required by pop and may be deleted
func (s Sightings) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Sighting) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Sighting) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Sighting) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
