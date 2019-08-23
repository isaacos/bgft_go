package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/isaac/project/bgft_go/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
