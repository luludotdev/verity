package handler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// RunBot starts the Discord bot listener
func RunBot() error {
	token := viper.GetString("token")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	session.AddHandler(onReady)

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
	user, _ := s.User("@me")
	fmt.Printf("Logged in as %s\n", user)
}
