package main

import (
	"fmt"
	"net/http"
	"encoding/json"
    "io/ioutil"
)

type Response struct {
	ResponseData ResponseData `json:"responseData"`
}

type ResponseData struct {
	TranslatedText string `json:"translatedText"`
}

func Translate(word string, sourceLang string, targetLang string) (string, error) {
	translationUrl := createUrl(word, sourceLang, targetLang)
	translatedWord, err := callApiToTranslate(translationUrl)
	if err != nil {
		return "", err
	}
	if translatedWord == word {
		return "", nil
	}
	reversedUrl := createUrl(word, targetLang, sourceLang)
	reverseTranslatedWord, err := callApiToTranslate(reversedUrl)
	if err != nil {
		return "", err
	}
	if reverseTranslatedWord != word {
		return "", err
	}
	return translatedWord, nil
}

func callApiToTranslate(url string) (string, error) {
	translationResponse, err := http.Get(url)
	if err != nil {
		return "", nil
	}

	responseBytes, err := ioutil.ReadAll(translationResponse.Body)
    if err != nil {
        return "", nil
    }

	var responseObject Response
	json.Unmarshal(responseBytes, &responseObject)

	translatedWord := responseObject.ResponseData.TranslatedText
	return translatedWord, nil
}

func createUrl(word string, sourceLang string, targetLang string) string {
	return fmt.Sprintf("https://api.mymemory.translated.net/get?q=%s&langpair=%s|%s", word, sourceLang, targetLang)
}