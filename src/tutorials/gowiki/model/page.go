package model

import (
	"os"

	"example.com/tutorial/gowiki/globais"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	// O literal octal 0600 indica que o arquivo deve ser
	// criado com permissões de leitura e gravação somente
	// para o usuário atual.
	return os.WriteFile(globais.PathToSave+filename, p.Body, 0600)
}
