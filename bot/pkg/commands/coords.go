package commands

import (
	"fmt"

	"github.com/lolPants/dgc"
	"github.com/lolPants/edsm-go"
)

func init() {
	router.RegisterCmd(&dgc.Command{
		Name:        "coords",
		Description: "Lookup XYZ coordinates for a given system",
		Usage:       "coords [system]",

		Handler: func(ctx *dgc.Ctx) {
			if ctx.Arguments.Amount() == 0 {
				ctx.RespondText("You must specify a system name!")
				return
			}

			arg := ctx.Arguments.AsSingle()
			systemName := arg.Raw()

			system, err := api.System(systemName, &edsm.SystemOptions{
				ShowCoordinates: true,
			})

			if err != nil {
				ctx.RespondText(fmt.Sprintf("Could not find a system named `%s`", systemName))
				return
			}

			if system.Coordinates == nil {
				ctx.RespondText(fmt.Sprintf("An error occured fetching coordinates for `%s`", system.Name))
				return
			}

			ctx.RespondText(fmt.Sprintf("```%+v```", system.Coordinates))
		},
	})
}
