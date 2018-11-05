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

type Sponsorship struct {
	ID                  uuid.UUID    `json:"id" db:"id"`
	CreatedAt           time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at" db:"updated_at"`
	EventID             uuid.UUID    `json:"event_id" db:"event_id"`
	UserID              uuid.UUID    `json:"user_id" db:"user_id"`
	UpdatedBy           uuid.UUID    `json:"updated_by" db:"updated_by"`
	Status              string       `json:"status" db:"status"`
	Type                nulls.String `json:"type" db:"type"`
	Benefits            nulls.String `json:"benefits" db:"benefits"`
	Costs               nulls.Int    `json:"costs" db:"costs"`
	BudgetQuarter       nulls.String `json:"budget_quarter" db:"budget_quarter"`
	SponsorContactName  nulls.String `json:"sponsor_contact_name" db:"sponsor_contact_name"`
	SponsorContactEmail nulls.String `json:"sponsor_contact_email" db:"sponsor_contact_email"`
	SponsorContactPhone nulls.String `json:"sponsor_contact_phone" db:"sponsor_contact_phone"`
	Staff               nulls.String `json:"staff" db:"staff"`
	RegistrationCode    nulls.String `json:"registration_code" db:"registration_code"`
	Notes               nulls.String `json:"notes" db:"notes"`
}

// String is not required by pop and may be deleted
func (s Sponsorship) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Sponsorships is not required by pop and may be deleted
type Sponsorships []Sponsorship

// String is not required by pop and may be deleted
func (s Sponsorships) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Sponsorship) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Sponsorship) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Sponsorship) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
