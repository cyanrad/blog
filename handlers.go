package main

import (
	"context"
	"io"
	"main/post"
	"main/template"
	"net/http"
	"strings"

	"github.com/a-h/templ"
)

// this is a terrible pattern
type Handler struct {
	pdb post.DB
}

func (h *Handler) handleBody() *templ.ComponentHandler {
	cards := generateCards(h.pdb.Posts)
	body := template.Body(cards)
	return templ.Handler(body)
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "methods not implemented", http.StatusNotFound)
		return
	}

	// getting post id
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		http.Error(w, "invalid URL path", http.StatusBadRequest)
		return
	}
	postId := pathSegments[2]

	// getting content
	// this is inefficient (getpost is running twice)
	postData, _ := h.pdb.GetPostById(postId)
	bytes, err := h.pdb.GetPostContent(postId)
	if err != nil {
		http.Error(w, "invalid URL path", http.StatusBadRequest)
		return
	}

	// creating post page
	post := generatePost(bytes, postData)
	post.Render(r.Context(), w)
}

func (h *Handler) handleFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, h.pdb.StaticPath+"favicon.ico")
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}
