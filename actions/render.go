package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
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
		},
	})
}
