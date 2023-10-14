package main

import (
	"context"
	"log"
	"main/templ"
	"net/http"
	"os"
)

type Person struct {
	Name string
}

func main() {
	basicCard := templ.Card("How I Made This Site", "devops", "4 Months Ago", "10 Minutes")
	body := templ.Body(basicCard)

	// this should relly be an env file
	indexFile, err := os.OpenFile("./static/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	body.Render(context.Background(), indexFile)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}
