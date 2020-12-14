package commands

import (
	"github.com/lolPants/dgc"
)

var (
	router *dgc.Router = dgc.Create(&dgc.Router{
		BotsAllowed:      false,
		IgnorePrefixCase: false,
		Prefixes:         []string{"v!"},
	})
)

// Router gets the pointer to the current command router
func Router() *dgc.Router {
	return router
}
