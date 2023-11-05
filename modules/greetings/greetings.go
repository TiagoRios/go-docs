package greetings

import (
	"errors"
	"fmt"
)

/*
Não sei se existe comentário de documentação. existe ??

Função retorna dois valores. isso abre muitas possibilidades.

Funções com nome em maiúsculo são exportadas.
*/
func Hello(name string) (string, error) {

	if name == "" { // não possui parênteses
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	
	// tipo nil(nada) alguma correlação com undefined do javascript?
	return message, nil
}
