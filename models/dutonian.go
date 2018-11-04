package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Dutonian struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	UserID    uuid.UUID    `json:"user_id" db:"user_id"`
	UpdatedBy uuid.UUID    `json:"updated_by" db:"updated_by"`
	Firstname string       `json:"firstname" db:"firstname"`
	Lastname  string       `json:"lastname" db:"lastname"`
	Bio       nulls.String `json:"bio" db:"bio"`
	Twitter   nulls.String `json:"twitter" db:"twitter"`
	Github    nulls.String `json:"github" db:"github"`
	Homepage  nulls.String `json:"homepage" db:"homepage"`
}

// String is not required by pop and may be deleted
func (d Dutonian) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Dutonians is not required by pop and may be deleted
type Dutonians []Dutonian

// String is not required by pop and may be deleted
func (d Dutonians) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Dutonian) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.Firstname, Name: "Firstname"},
		&validators.StringIsPresent{Field: d.Lastname, Name: "Lastname"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Dutonian) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Dutonian) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
