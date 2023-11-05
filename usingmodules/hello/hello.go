package main

import (
	"fmt"

	"example.com/greetings" // minha própria dependência.
)

func main() {
	message := greetings.Hello("DEVmon")
	fmt.Println(message)
}
