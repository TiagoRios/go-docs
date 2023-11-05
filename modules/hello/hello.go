package main

import (
	"fmt"
	"log"

	"example.com/greetings" // minha própria dependência.
)

func main() {
	log.SetPrefix("greetings: ") // Prefixo dos logs
	log.SetFlags(0)              // Disable printing the time, source file, and line number

	message, err := greetings.Hello("") // gera um erro 
	// message, err := greetings.Hello("DEVmon")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
