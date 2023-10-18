package main

import (
	"main/template"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	basicCard := template.Card("How I Made This Site", "devops", "4 Months Ago", "10 Minutes")
	body := template.Body(basicCard)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", templ.Handler(body))

	http.ListenAndServe(":3000", nil)
}
