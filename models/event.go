package models

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
)

type Event struct {
	ID               uuid.UUID    `json:"id" db:"id"`
	CreatedAt        time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at" db:"updated_at"`
	UserID           uuid.UUID    `json:"user_id" db:"user_id"`
	UpdatedBy        uuid.UUID    `json:"updated_by" db:"updated_by"`
	Title            string       `json:"title" db:"title"`
	EventBeginDate   time.Time    `json:"event_begin_date" db:"event_begin_date"`
	EventEndDate     time.Time    `json:"event_end_date" db:"event_end_date"`
	CfpBeginDate     time.Time    `json:"cfp_begin_date" db:"cfp_begin_date"`
	CfpEndDate       time.Time    `json:"cfp_end_date" db:"cfp_end_date"`
	Location         nulls.String `json:"location" db:"location"`
	HomePage         nulls.String `json:"home_page" db:"home_page"`
	SchedulePage     nulls.String `json:"schedule_page" db:"schedule_page"`
	SponsorPage      nulls.String `json:"sponsor_page" db:"sponsor_page"`
	CfpPage          nulls.String `json:"cfp_page" db:"cfp_page"`
	RegistrationPage nulls.String `json:"registration_page" db:"registration_page"`
	EventReport      nulls.String `json:"event_report" db:"event_report"`
	Attendance       nulls.Int    `json:"attendance" db:"attendance"`
	Logo             binding.File `db:"-"`
	LogoName         string       `json:"logo_name" db:"logo_name"`
	Proposals        Proposals    `has_many:"proposals" order_by:"created_at asc"`
	Sponsorships     Sponsorships `has_many:"sponsorships" order_by:"created_at asc"`
}

// SelectLabel - label for select tag options
func (e Event) SelectLabel() string {
	return e.Title
}

// SelectValue - value for select tag options
func (e Event) SelectValue() interface{} {
	return e.ID
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
func (e *Event) AfterCreate(tx *pop.Connection) error {
	if !e.Logo.Valid() {
		return nil
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(os.Getenv("S3_REGION"))
		log.Fatal(err)
	}

	dir := filepath.Join(".", "uploads", "events")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	// TODO: Create a function using SHA and a salt for the new filename and use that instead of e.Logo.Filename below.
	f, err := os.Create(filepath.Join(dir, e.Logo.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, e.Logo)

	// Upload
	err = AddFileToS3(s, filepath.Join(dir, e.Logo.Filename))
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (e *Event) AfterUpdate(tx *pop.Connection) error {
	if !e.Logo.Valid() {
		return nil
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(os.Getenv("S3_REGION"))
		log.Fatal(err)
	}

	dir := filepath.Join(".", "uploads", "events")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, e.Logo.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, e.Logo)

	// Upload
	err = AddFileToS3(s, filepath.Join(dir, e.Logo.Filename))
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Title, Name: "Title"},
		&validators.TimeIsBeforeTime{
			FirstName:  "Event Start Date",
			FirstTime:  e.EventBeginDate,
			SecondName: "Event End Date",
			SecondTime: e.EventEndDate,
		},
		&validators.TimeIsBeforeTime{
			FirstName:  "CFP Begin Date",
			FirstTime:  e.CfpBeginDate,
			SecondName: "CFP End Date",
			SecondTime: e.CfpEndDate,
		},
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
