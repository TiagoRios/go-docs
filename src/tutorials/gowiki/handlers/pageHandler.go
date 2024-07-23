package handlers

import (
	"fmt"
	"net/http"

	"example.com/tutorial/gowiki/globais"
	"example.com/tutorial/gowiki/model"
	"example.com/tutorial/gowiki/service"
)

func Handler(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hi there, I love %s!", r.URL.Path[1:])
}

func ViewHandler(response http.ResponseWriter, r *http.Request, title string) {
	page, err := service.LoadPage(title)

	// não encontrou cria novo arquivo
	if err != nil {
		http.Redirect(response, r, "/edit/"+title, http.StatusFound)
		return
	}

	RenderTemplate(response, "view", page)
}

func EditHandler(response http.ResponseWriter, r *http.Request, title string) {
	page, err := service.LoadPage(title)

	if err != nil {
		page = &model.Page{Title: title}
	}

	RenderTemplate(response, "edit", page)
}

func SaveHandler(response http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	page := &model.Page{Title: title, Body: []byte(body)}

	err := page.Save()

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(response, r, "/view/"+title, http.StatusFound)
}

func RenderTemplate(response http.ResponseWriter, templateName string, page *model.Page) {
	// vá ao template XXX.html e renderize as informaçoes do struct Page no template
	err := globais.Templates.ExecuteTemplate(response, templateName+".html", page)

	if err != nil { // houve erro
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
