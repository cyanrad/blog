package main

import (
	"log"
	"main/post"
	"main/template"
	"net/http"
	"path"
	"strings"

	"github.com/a-h/templ"
)

func main() {
	// this should really be an env file
	p, err := post.LoadPostsData("posts.json")
	if err != nil {
		log.Fatal(err)
	}

	// idk why but i really hate how I'm organizing code right now
	relativeTime, err := post.GetRelativePastTime(p[0].ReleaseDate)
	if err != nil {
		log.Fatal(err)
	}

	basicCard := template.Card(p[0].Title,
		p[0].Topic,
		relativeTime,
		post.GetReadingTime(p[0].WordCount))

	body := template.Body(basicCard)

	http.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)
	http.HandleFunc("/post/", handlePost)
	http.Handle("/", templ.Handler(body))

	http.ListenAndServe(":3000", nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cleanedPath := path.Clean(r.URL.Path)
	cleanedPath = cleanedPath[1:] // Remove the leading '/'
	segments := strings.Split(cleanedPath, "/")

	http.ServeFile(w, r, "./static/posts/"+segments[1])
}
