package handler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"

	"github.com/lolPants/verity/bot/pkg/commands"
)

// RunBot starts the Discord bot listener
func RunBot() error {
	token := viper.GetString("token")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	session.AddHandler(onReady)
	router := commands.Router()
	router.RegisterDefaultHelpCommand(session, nil)
	router.Initialize(session)

	fmt.Println("Logging in to Discord...")
	err = session.Open()
	if err != nil {
		return err
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Shutting down...")
	session.Close()

	return nil
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Logged in as %s\n", s.State.User)
	s.UpdateStatus(0, "v!help")
}
