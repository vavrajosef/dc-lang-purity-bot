package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Hello command",
		},
	}
	commandHandlers = map[string]func(dcSession *discordgo.Session, dcInteraction *discordgo.InteractionCreate){
		"hello": func(dcSess *discordgo.Session, dcInter *discordgo.InteractionCreate) {
			dcSess.InteractionRespond(dcInter.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello",
				},
			})
		},
	}
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(dcSession *discordgo.Session, dcInter *discordgo.InteractionCreate) {
		if handlerVar, ok := commandHandlers[dcInter.ApplicationCommandData().Name]; ok {
			handlerVar(dcSession, dcInter)
		}
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	log.Println("Adding commands...")
	for _, commandValue := range commands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", commandValue)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", commandValue.Name, err)
		}
	}

	defer dg.Close()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
