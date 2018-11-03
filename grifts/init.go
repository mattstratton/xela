package grifts

import (
	"github.com/PagerDuty/xela/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
