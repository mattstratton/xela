package actions

import (
	"fmt"
	"os"

	"github.com/PagerDuty/xela/models"
	"github.com/gobuffalo/buffalo"
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
			"getUserEmail": func(c buffalo.Context) string {

				// // Get the DB connection from the context
				// tx := c.Value("tx").(*pop.Connection)
				// user_id := "c2fbcdfc-3b19-410d-bd43-a2f18b6b309e"
				// // Allocate an empty User
				// user := &models.User{}

				// // To find the User the parameter user_id is used.
				// err := tx.Find(&user, user_id)

				// if err != nil {
				// 	return "error"
				// } else {

				// 	// return user.Email.String
				// 	return "mstratton@pagerduty.com"
				// }
				return "mstratton@pagerduty.com"
			},
			"getSpeakerName": func(dutonianID uuid.UUID, c plush.HelperContext) (string, error) {
				// Get the DB connection from the context
				tx, ok := c.Value("tx").(*pop.Connection)
				if !ok {
					return "", errors.WithStack(errors.New("no transaction found"))
				}

				dutonian := &models.Dutonian{}

				err := tx.Find(&dutonian, dutonianID)
				if err != nil {
					return "", errors.WithStack(errors.New("query error"))
				} else {
					return "Matt Stratton", nil
				}

			},
			"buildS3Url": func() string {
				return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/", os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET"))
			},
		},
	})
}
