package main

import (
	"log"
	"net/http"

	handlers "example.com/tutorial/gowiki/handlers"
	"example.com/tutorial/gowiki/model"
	"example.com/tutorial/gowiki/util"
)

func HomeHandler(response http.ResponseWriter, r *http.Request) {
	page := &model.Page{Title: "index"}

	// renderize o template com os dados do struct page
	handlers.RenderTemplate(response, "index", page)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/view/", util.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", util.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", util.MakeHandler(handlers.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
