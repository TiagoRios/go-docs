package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
Não sei se existe comentário de documentação. existe ??

Função retorna dois valores. isso abre muitas possibilidades.

Funções com nome em maiúsculo são exportadas.
*/
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Hail, %v! Well met!",
		"Great to see you, %v!",
		"Olá %v, seja bem vindo!",
		"Hola %v, see bienvenido!",
		"привет %v, добро пожаловать",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
