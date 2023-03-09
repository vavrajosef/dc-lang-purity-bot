package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

var TRANSLATE_REGEX = regexp.MustCompile(`^translate\s[a-z]+\s[a-z]+$`)

func main() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents |= discordgo.IntentMessageContent
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if TRANSLATE_REGEX.MatchString(m.Content) {
		fmt.Println("Translating ...")
		if m.MessageReference == nil {
			fmt.Println("No previous message to translate")
		} else {
			messageReferenceContent, err := GetMessageReferenceContent(s, m.MessageReference.ChannelID, m.MessageReference.MessageID)
			if err != nil {
				fmt.Println("Error occured: " + err.Error())
				return
			}
			translatedMessage, err := ProcessMessage(m.Content, messageReferenceContent)
			if err != nil {
				fmt.Println("Error occured: " + err.Error())
			} else {
				fmt.Println("Translation success: this is result " + translatedMessage)
				if translatedMessage == "" {
					s.ChannelMessageSend(m.ChannelID, "Nothing to translate, text seems pure.")
				} else {
					s.ChannelMessageSend(m.ChannelID, translatedMessage)
				}
			}
		}
	}
}