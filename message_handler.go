package main

import (
	"fmt"
	"strings"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/language"
)

var IS_WORD_REGEX = regexp.MustCompile(`^[a-zA-Z]+[,.]?$`)

func ProcessMessage(translationInstructions string, previousMessage string) (string, error) {
	instructionSlice := strings.Split(translationInstructions, " ")
	instructionsLength := len(instructionSlice)
	if instructionsLength != 3 {
		return "", fmt.Errorf("illegal number of instructions. %d is not supported: %s", instructionsLength, translationInstructions)
	}
	sourceLanguageString := instructionSlice[1]
	targetLanguageString := instructionSlice[2]
	err := CheckLanguage(sourceLanguageString)
	if err != nil {
		return "", err
	}
	err = CheckLanguage(targetLanguageString)
	if err != nil {
		return "", err
	}
	result := make([]string, 0)
	for _, word := range strings.Split(previousMessage, " ") {
		if IS_WORD_REGEX.MatchString(word) {
			translatedWord, err := Translate(word, sourceLanguageString, targetLanguageString)
			if err != nil {
				return "", err
			}
			if translatedWord != "" {
				result = append(result, translatedWord)
			}
		}
	}
	if len(result) == 0 {
		return "", nil
	}
	return "*" + strings.Join(result, ", "), nil
}


func CheckLanguage(lang string) error {
	_, err := language.Parse(lang)
	return err
}

func GetMessageReferenceContent(session *discordgo.Session, channelId string, messageId string) (string, error) {
	message, err := session.ChannelMessage(channelId, messageId)
	if err != nil {
		return "", err
	}
	return message.Content, nil
}