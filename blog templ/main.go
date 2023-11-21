package main

import (
	"context"
	"io"
	"log"
	"main/post"
	"main/template"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"
)

var POSTS []post.Post

const POSTS_PATH = "./static/posts/"
const POSTS_CONFIG_FILE = "posts.json"

func main() {
	var err error
	POSTS, err = post.LoadPostsData(POSTS_CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	cards := generateCards()
	body := template.Body(cards)

	// this is absolutely terrible but I'm fine with it now
	http.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)
	http.HandleFunc("/post/", handlePost)
	http.Handle("/", templ.Handler(body))

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func generateCards() []templ.Component {
	cards := make([]templ.Component, len(POSTS))

	for i, p := range POSTS {
		relativeDuration, err := post.GetRelativeDuration(p.ReleaseDate)
		if err != nil {
			log.Fatal(err)
		}

		cards[i] = template.Card(
			p.Title,
			p.Id,
			p.Topic,
			relativeDuration,
			post.GetReadingTime(p.WordCount),
		)
	}

	return cards
}

// format: /post/%d
func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "methods not implemented", http.StatusNotFound)
		return
	}

	pathSegments := strings.Split(r.URL.Path, "/")

	// finding post
	var postData post.Post
	found := false
	for _, p := range POSTS {
		if p.Id == pathSegments[2] {
			postData = p
			found = true
		}
	}
	if !found {
		http.Error(w, "invalid URL path", http.StatusBadRequest)
		return
	}

	// reading post file context
	postPath := POSTS_PATH + postData.Title + ".html"
	f, err := os.Open(postPath)
	if err != nil {
		log.Printf("post %s was found in posts config but not at %s\n", postData.Id, postPath)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Println("reading existing file failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// creating post page
	styledHTML := post.StyleHTML(bytes)

	relativeTime, err := post.GetRelativeDuration(postData.ReleaseDate)
	if err != nil {
		log.Fatal(err)
	}

	post := template.Post(
		Unsafe(string(styledHTML)),
		postData.Title,
		postData.Hook,
		postData.Topic,
		relativeTime,
		post.GetReadingTime(postData.WordCount),
	)

	post.Render(r.Context(), w)
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}
