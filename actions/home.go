package actions

import (
	"fmt"
	"time"

	"github.com/PagerDuty/xela/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	events := &models.Events{}

	qs := fmt.Sprintf("event_begin_date >= '%s'", time.Now().Format("2006-01-02 00:00:00"))
	// q := tx.Where("event_begin_date >= '11/01/2018'")
	q := tx.Where(qs)

	// 2006-01-02 00:00:00
	// Mon Jan 2 15:04:05 MST 2006

	// q := tx.PaginateFromParams(c.Request().URL.Query())
	// You can order your list here. Just change
	err := q.Order("event_begin_date asc").All(events)
	// to:
	// err := tx.Order("create_at desc").All(events)
	if err != nil {
		return errors.WithStack(err)
	}

	// Make events available inside the html template
	c.Set("events", events)
	// c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("index.html"))
}
