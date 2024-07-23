package globais

import (
	"regexp"
	"text/template"
)

var PathToSave = "data/"
var PathToFind = "template/"

// Define um padrão de URL permitida.
var ValidPath = regexp.MustCompile("^/(edit|save|view|)/([a-zA-Z0-9]+)$")

// whitespace = [\t\n\f\r]

var Templates = template.Must(template.ParseFiles(
	PathToFind+"edit.html",
	PathToFind+"view.html",
	PathToFind+"index.html"))
