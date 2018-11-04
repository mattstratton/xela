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

type Abstract struct {
	ID                     uuid.UUID    `json:"id" db:"id"`
	CreatedAt              time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at" db:"updated_at"`
	UserID                 uuid.UUID    `json:"user_id" db:"user_id"`
	UpdatedBy              uuid.UUID    `json:"updated_by" db:"updated_by"`
	Title                  string       `json:"title" db:"title"`
	OneSentenceDescription nulls.String `json:"one_sentence_description" db:"one_sentence_description"`
	ShortDescription       nulls.String `json:"short_description" db:"short_description"`
	Abstract               string       `json:"abstract" db:"abstract"`
	TalkType               nulls.String `json:"talk_type" db:"talk_type"`
}

// String is not required by pop and may be deleted
func (a Abstract) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Abstracts is not required by pop and may be deleted
type Abstracts []Abstract

// String is not required by pop and may be deleted
func (a Abstracts) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Abstract) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Title, Name: "Title"},
		&validators.StringIsPresent{Field: a.Abstract, Name: "Abstract"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Abstract) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Abstract) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
