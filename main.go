package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("quasarify: ")
	// message := "this is a normal message from earth";
	quasar, err := QuasarifyWords([]string{"What's", "up", "dawg?"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quasar)
}
