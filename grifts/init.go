package grifts

import (
	"github.com/devrelcollective/xela/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
