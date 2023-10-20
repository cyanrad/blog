package main

import (
	"context"
	"io"
	"log"
	"main/post"
	"main/template"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/a-h/templ"
)

var POSTS []post.Post

func main() {
	// this should really be an env file
	var err error
	POSTS, err = post.LoadPostsData("posts.json")
	if err != nil {
		log.Fatal(err)
	}

	// idk why but i really hate how I'm organizing code right now
	relativeTime, err := post.GetRelativePastTime(POSTS[0].ReleaseDate)
	if err != nil {
		log.Fatal(err)
	}

	basicCard := template.Card(POSTS[0].Title,
		POSTS[0].Topic,
		relativeTime,
		post.GetReadingTime(POSTS[0].WordCount))

	body := template.Body(basicCard)

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

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "methods not implemented", http.StatusNotFound)
		return
	}

	cleanedPath := path.Clean(r.URL.Path)
	cleanedPath = cleanedPath[1:] // Remove the leading '/'
	pathSegments := strings.Split(cleanedPath, "/")

	// should probably check for .html at the end

	filePath, err := url.PathUnescape(pathSegments[1])
	if err != nil {
		log.Println("failed to unescape url path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postPath := "./static/posts/" + filePath
	if containsDotDot(postPath) {
		http.Error(w, "invalid URL path", http.StatusBadRequest)
		return
	}

	f, err := os.Open(postPath)
	if err != nil {
		http.Error(w, "invalid URL path", http.StatusBadRequest)
		return
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Println("reading existing file failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	styledHTML := post.StyleHTML(bytes)

	// finding post
	var postData post.Post
	found := false
	for _, p := range POSTS {
		// removing the .html
		if filePath[:len(filePath)-5] == p.Title {
			postData = p
			found = true
		}
	}
	if !found {
		// this should be an env file
		log.Println("post was found in /static/posts file but was not found in posts config data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	relativeTime, err := post.GetRelativePastTime(postData.ReleaseDate)
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

// taken form the http standard package in fs.go
func containsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, isSlashRune) {
		if ent == ".." {
			return true
		}
	}
	return false
}

// taken form the http standard package in fs.go
func isSlashRune(r rune) bool { return r == '/' || r == '\\' }
