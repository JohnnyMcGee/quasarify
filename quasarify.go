package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Quasarify(message string) (string, error) {
	if message == "" {
		return "", errors.New("empty message")
	}
	newMessage := fmt.Sprintf(randomFormat(), message)
	return newMessage, nil
}

func QuasarifyWords(words []string) (string, error) {
	qresponse := "👽"

	for _, word := range words {
		translation, err := Quasarify(word)
		if err != nil {
			return "", err
		}
		qresponse += "\n"
		qresponse += translation
	}
	qresponse += "\n🚀"
	return qresponse, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Greetings, %v. I come in peace!",
		"Quasar transmission: %v... in Space!",
		"00111011011Computer Transimission: %v. --sent from the Virgo A. galaxy000111111111110110101001111010111",
	}
	return formats[rand.Intn(len(formats))]
}
