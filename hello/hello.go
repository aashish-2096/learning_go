package main

import (
	"fmt"
	"log"

	"learning.go/greetings"
)

func main() {
	names := []string{"Aashish", "Kumar", "Legend"}
	// message, err := greetings.Hello("Aashish")
	name, age := "Aashish", 23
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, age)
	name, age = "Tp", 34
	fmt.Println(name, age)
	fmt.Println(message)

	var namesT []string
	if namesT == nil {
		fmt.Println("Capacity1", cap(namesT))
		fmt.Println("slice is nil going to append")
		namesT = append(namesT, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", namesT)
		fmt.Println("Capacity2", cap(namesT))
	}
}
