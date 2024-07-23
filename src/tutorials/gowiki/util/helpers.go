package util

import (
	"errors"
	"net/http"

	"example.com/tutorial/gowiki/globais"
)

type myHandlerFunc func(http.ResponseWriter, *http.Request, string)

// substituiu a função GetTitle
func MakeHandler(fn myHandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		match := globais.ValidPath.FindStringSubmatch(request.URL.Path)

		if match == nil { // houve erro
			http.NotFound(response, request)
			return
		}

		// localhost:8080/
		// se a chamada for pela raiz não vai existir o index 2
		title := match[2]

		// saveHandler(response, request, title)
		fn(response, request, title) // chamando a função passada como argumento
	}
}

func GetTitle(response http.ResponseWriter, r *http.Request) (string, error) {
	match := globais.ValidPath.FindStringSubmatch(r.URL.Path)

	if match == nil {
		http.NotFound(response, r)
		return "", errors.New("invalid Page Title")
	}

	return match[2], nil // The title is the second subexpression.
}
