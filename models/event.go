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

type Event struct {
	ID               uuid.UUID    `json:"id" db:"id"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" db:"updated_at"`
	UserID           uuid.UUID    `json:"user_id" db:"user_id"`
	UpdatedBy        uuid.UUID    `json:"updated_by" db:"updated_by"`
	Title            string       `json:"title" db:"title"`
	Location         nulls.String `json:"location" db:"location"`
	HomePage         nulls.String `json:"home_page" db:"home_page"`
	SchedulePage     nulls.String `json:"schedule_page" db:"schedule_page"`
	SponsorPage      nulls.String `json:"sponsor_page" db:"sponsor_page"`
	CfpPage          nulls.String `json:"cfp_page" db:"cfp_page"`
	RegistrationPage nulls.String `json:"registration_page" db:"registration_page"`
	EventReport      nulls.String `json:"event_report" db:"event_report"`
	Attendance       nulls.Int    `json:"attendance" db:"attendance"`
}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Title, Name: "Title"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
