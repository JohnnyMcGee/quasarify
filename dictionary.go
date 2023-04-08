package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func isNoun(word string) (bool, error) {
	wordLookup, err := DictionaryLookup(word)
	notFound := len(wordLookup.Meanings) == 0
	if err != nil || notFound {
		return false, err
	}
	pos := wordLookup.Meanings[0].PartOfSpeech
	return pos == "noun", nil
}

func DictionaryLookup(word string) (Word, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%v", word)
	resp, err := http.Get(url)
	if err != nil {
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	}

	var responseObject Response
	json.Unmarshal(body, &responseObject)

	if len(responseObject) == 0 {
		return Word{Meanings: []Meaning{}}, nil
	}
	return responseObject[0], nil
}

type Response = []Word

type Word struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string `json:"partOfSpeech"`
}
