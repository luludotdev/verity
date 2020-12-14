package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lolPants/dgc"
)

func init() {
	permissions := discordgo.PermissionSendMessages |
		discordgo.PermissionEmbedLinks |
		discordgo.PermissionAttachFiles |
		discordgo.PermissionReadMessageHistory |
		discordgo.PermissionAddReactions

	router.RegisterCmd(&dgc.Command{
		Name:        "invite",
		Description: "Generates an invite link for Verity.",

		Handler: func(ctx *dgc.Ctx) {
			invite := fmt.Sprintf("<https://discord.com/api/oauth2/authorize?client_id=%s&scope=bot&permissions=%d>", ctx.Session.State.User.ID, permissions)
			ctx.RespondText(invite)
		},
	})
}
