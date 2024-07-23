package service

import (
	"os"

	"example.com/tutorial/gowiki/model"
)

var PathToSave = "data/"

// carrega dados da pagina. não carrega a pagina.
func LoadPage(title string) (*model.Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(PathToSave + filename)

	if err != nil {
		return nil, err
	}

	return &model.Page{Title: title, Body: body}, nil
}
