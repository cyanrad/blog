package main

import (
	"log"
	"main/post"
	"main/template"
	"main/util"

	"github.com/a-h/templ"
)

func generateCards(posts []post.Post) []templ.Component {
	cards := make([]templ.Component, len(posts))

	for i, p := range posts {
		relativeDuration, err := util.GetRelativeDuration(p.ReleaseDate)
		if err != nil {
			log.Fatal(err)
		}

		cards[i] = template.Card(
			p.Title,
			p.Id,
			p.Topic,
			relativeDuration,
			util.GetReadingTime(p.WordCount),
		)
	}

	return cards
}

func generatePost(postContent []byte, postData post.Post) templ.Component {
	styledHTML := post.StyleHTML(postContent)

	relativeTime, err := util.GetRelativeDuration(postData.ReleaseDate)
	if err != nil {
		log.Fatal(err)
	}

	return template.Post(
		Unsafe(string(styledHTML)),
		postData.Title,
		postData.Hook,
		postData.Topic,
		relativeTime,
		util.GetReadingTime(postData.WordCount),
	)

}
