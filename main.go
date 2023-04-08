package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("quasarify: ")
	message := "this is a normal message from earth"
	quasar, err := QuasarifyMessage(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quasar)
}
