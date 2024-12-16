package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type Post struct {
	Id          string
	Title       string
	Hook        string
	Topic       string
	ReleaseDate time.Time
	WordCount   int
}

type DB struct {
	Posts      []Post
	ConfigPath string
	PostsPath  string
	StaticPath string
}

func (db *DB) LoadPostsConfig() error {
	f, err := os.Open(db.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to open posts config: %w", err)
	}

	posts, err := decodePostsJson(f)
	if err != nil {
		return fmt.Errorf("failed to decode posts config: %w", err)
	}

	db.Posts = posts
	return nil
}

func (db *DB) GetPostContent(id string) ([]byte, error) {
	post, ok := db.GetPostById(id)
	if !ok {
		return nil, fmt.Errorf("post id %s doesn't exist", id)
	}

	postPath := db.PostsPath + post.Title + ".html"
	f, err := os.Open(postPath)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("unexpected error: post %s was found in posts config but failed to open", id), err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("unexpected error: reading existing post %s failed", id), err)
	}

	return bytes, nil
}

func (db *DB) GetPostById(id string) (Post, bool) {
	var post Post
	found := false

	for _, p := range db.Posts {
		if p.Id == id {
			post = p
			found = true
		}
	}

	return post, found
}

func decodePostsJson(r io.Reader) ([]Post, error) {
	d := json.NewDecoder(r)
	var posts []Post

	err := d.Decode(&posts)
	if err != nil {
		return nil, fmt.Errorf("failed to decode posts json: %w", err)
	}

	return posts, nil
}
