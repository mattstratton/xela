package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Proposal struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	UpdatedBy  uuid.UUID `json:"updated_by" db:"updated_by"`
	SpeakerID  uuid.UUID `json:"speaker_id" db:"speaker_id"`
	AbstractID uuid.UUID `json:"abstract_id" db:"abstract_id"`
	EventID    uuid.UUID `json:"event_id" db:"event_id"`
	Status     string    `json:"status" db:"status"`
}

// String is not required by pop and may be deleted
func (p Proposal) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Proposals is not required by pop and may be deleted
type Proposals []Proposal

// String is not required by pop and may be deleted
func (p Proposals) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Proposal) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Proposal) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Proposal) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
