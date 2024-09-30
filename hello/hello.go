package main

import (
	"fmt"

	"log"

	"learning.go/greetings"
)

func main() {
	names := []string{"Aashish", "Kumar", "Legend"}
	// message, err := greetings.Hello("Aashish")
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
