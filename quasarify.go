package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func QuasarifyMessage(message string) (string, error) {
	words := strings.Fields(message)
	qwords, err := QuasarifyWords(words)
	if err != nil {
		return message, err
	}
	qmessage := strings.Join(qwords, " ")
	return qmessage, nil
}

func QuasarifyWords(words []string) ([]string, error) {
	qwords := []string{}

	for _, word := range words {
		translation, err := Quasarify(word)
		if err != nil {
			return []string{}, err
		}
		qwords = append(qwords, translation)
	}
	return qwords, nil
}

func Quasarify(word string) (string, error) {
	if word == "" {
		return "", nil
	}
	isNoun, err := isNoun(word)
	if err != nil {
		return word, err
	}
	if !isNoun {
		return word, nil
	}
	newWord := fmt.Sprintf(randomFormat(), word)
	return newWord, nil
}

func randomFormat() string {
	formats := []string{
		"space-%v",
		"galactic-%v",
		"light-speed-%v",
		"proton-%v",
		"quasar-%v",
	}
	return formats[rand.Intn(len(formats))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
