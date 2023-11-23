package main

import (
	"log"
	"main/post"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("BLOG_PORT")
	URL := os.Getenv("BLOG_URL")
	POSTS_CONFIG_PATH := os.Getenv("BLOG_POSTS_CONFIG_PATH")
	CONTENT_PATH := os.Getenv("BLOG_CONTENT_PATH")
	POSTS_PATH := CONTENT_PATH + "posts/"
	STATIC_PATH := CONTENT_PATH + "static/"

	// loading data
	pdb := post.DB{
		Posts:      nil,
		ConfigPath: POSTS_CONFIG_PATH,
		PostsPath:  POSTS_PATH,
	}
	err = pdb.LoadPostsConfig()
	if err != nil {
		log.Fatal(err)
	}

	// endpoint handlers
	h := Handler{pdb}
	http.Handle("/", h.handleBody())
	http.HandleFunc("/post/", h.handlePost)
	http.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(STATIC_PATH))),
	)

	// starting server
	err = http.ListenAndServe(URL+":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
