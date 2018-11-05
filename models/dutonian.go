package models

import (
	"encoding/json"
	"fmt"
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
	Photo     binding.File `db:"-"`
	PhotoName string       `json:"photo_name" db:"photo_name"`
}

// SelectLabel - label for select tag options
func (d Dutonian) SelectLabel() string {
	return fmt.Sprintf("%s %s", d.Firstname, d.Lastname)
}

// SelectValue - value for select tag options
func (d Dutonian) SelectValue() interface{} {
	return d.ID
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

func (d *Dutonian) AfterCreate(tx *pop.Connection) error {
	if !d.Photo.Valid() {
		return nil
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(os.Getenv("S3_REGION"))
		log.Fatal(err)
	}

	dir := filepath.Join(".", "uploads", "dutonians", d.Lastname, d.Firstname)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, d.Photo.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, d.Photo)

	// Upload
	err = AddFileToS3(s, filepath.Join(dir, d.Photo.Filename))
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (d *Dutonian) AfterUpdate(tx *pop.Connection) error {
	if !d.Photo.Valid() {
		return nil
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("S3_REGION"))})
	if err != nil {
		log.Fatal(os.Getenv("S3_REGION"))
		log.Fatal(err)
	}

	dir := filepath.Join(".", "uploads", "dutonians", d.Lastname, d.Firstname)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, d.Photo.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, d.Photo)

	// Upload
	err = AddFileToS3(s, filepath.Join(dir, d.Photo.Filename))
	if err != nil {
		log.Fatal(err)
	}
	return err
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
