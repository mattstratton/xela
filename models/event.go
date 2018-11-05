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
	EventBeginDate   nulls.Time   `json:"event_begin_date" db:"event_begin_date"`
	EventEndDate     nulls.Time   `json:"event_end_date" db:"event_end_date"`
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

// func AddFileToS3(s *session.Session, fileDir string) error {

// 	// Open the file for use
// 	file, err := os.Open(fileDir)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Get file size and read the file content into a buffer
// 	fileInfo, _ := file.Stat()
// 	var size int64 = fileInfo.Size()
// 	buffer := make([]byte, size)
// 	file.Read(buffer)

// 	// Config settings: this is where you choose the bucket, filename, content-type etc.
// 	// of the file you're uploading.
// 	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
// 		Bucket:               aws.String(os.Getenv("S3_BUCKET")),
// 		Key:                  aws.String(fileDir),
// 		Body:                 bytes.NewReader(buffer),
// 		ACL:                  aws.String("public-read"),
// 		ContentLength:        aws.Int64(size),
// 		ContentType:          aws.String(http.DetectContentType(buffer)),
// 		ContentDisposition:   aws.String("attachment"),
// 		ServerSideEncryption: aws.String("AES256"),
// 	})
// 	return err
// }
