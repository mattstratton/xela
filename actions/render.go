package actions

import (
	"fmt"
	"os"
	"time"

	"github.com/devrelcollective/xela/models"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
			"checkNull": func() {

			},
			"getUserEmail": func(userID uuid.UUID, c plush.HelperContext) (string, error) {

				// Get the DB connection from the context
				tx, ok := c.Value("tx").(*pop.Connection)
				if !ok {
					return "", errors.WithStack(errors.New("no transaction found"))
				}
				// Allocate an empty User
				user := &models.User{}

				// To find the User the parameter user_id is used.
				err := tx.Find(user, userID)

				if err != nil {
					return "", errors.WithStack(errors.New("query error"))
				}
				return user.Email.String, nil
			},
			"getSpeakerName": func(dutonianID uuid.UUID, c plush.HelperContext) (string, error) {
				// Get the DB connection from the context
				tx, ok := c.Value("tx").(*pop.Connection)
				if !ok {
					return "", errors.WithStack(errors.New("no transaction found"))
				}

				dutonian := &models.Dutonian{}

				err := tx.Find(dutonian, dutonianID)
				if err != nil {
					return "", errors.WithStack(errors.New("query error"))
				} else {
					return fmt.Sprintf("%s %s", dutonian.Firstname, dutonian.Lastname), nil
				}

			},
			"buildS3Url": func() string {
				return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/", os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET"))
			},
			"getAbstractName": func(abstractID uuid.UUID, c plush.HelperContext) (string, error) {
				// Get the DB connection from the context
				tx, ok := c.Value("tx").(*pop.Connection)
				if !ok {
					return "", errors.WithStack(errors.New("no transaction found"))
				}

				abstract := &models.Abstract{}

				err := tx.Find(abstract, abstractID)
				if err != nil {
					return "", errors.WithStack(errors.New("query error"))
				} else {
					return abstract.Title, nil
				}

			},
			"getEventName": func(eventID uuid.UUID, c plush.HelperContext) (string, error) {
				// Get the DB connection from the context
				tx, ok := c.Value("tx").(*pop.Connection)
				if !ok {
					return "", errors.WithStack(errors.New("no transaction found"))
				}

				event := &models.Event{}

				err := tx.Find(event, eventID)
				if err != nil {
					return "", errors.WithStack(errors.New("query error"))
				} else {
					return event.Title, nil
				}
			},
			"getEventLocation": func(eventID uuid.UUID, c plush.HelperContext) (string, error) {
				// Get the DB connection from the context
				tx, ok := c.Value("tx").(*pop.Connection)
				if !ok {
					return "", errors.WithStack(errors.New("no transaction found"))
				}

				event := &models.Event{}

				err := tx.Find(event, eventID)
				if err != nil {
					return "", errors.WithStack(errors.New("query error"))
				} else {
					return event.Location.String, nil
				}
			},
			"displayEventDate": func(eventStartDate time.Time, eventEndDate time.Time) string {
				if eventStartDate == eventEndDate {
					return eventStartDate.Format("January 02, 2006")
				} else {
					return fmt.Sprintf("%s - %s", eventStartDate.Format("January 02, 2006"), eventEndDate.Format("January 02, 2006"))
				}
			},
		},
	})
}
