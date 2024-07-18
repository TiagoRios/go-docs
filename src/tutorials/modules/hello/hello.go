package main

import (
	"fmt"
	"log"

	"example.com/greetings" // minha própria dependência.
)

func main() {
	log.SetPrefix("greetings: ") // Prefixo dos logs
	log.SetFlags(0)              // Disable printing the time, source file, and line number

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(messages)
	fmt.Println()
}
