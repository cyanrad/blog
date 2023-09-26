package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string
}

func main() {
	basicCard := Card("/1", "internship at the emirates", "static/bg.png", "devops", "4 Months Ago", "10 Minutes")
	body := Body(basicCard)

	indexFile, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	body.Render(context.Background(), indexFile)

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":3000", nil)
}
