package post

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestDecodePostsJson(t *testing.T) {
	jsonString := `[
    {
		"id": "1",
        "title": "How I Made This Site",
        "hook": "In short Golang, Templ, and RAW DOGGING HTML with the power of ZEUS",
        "topic": "web dev",
        "releaseDate": "2023-10-17T00:00:00Z",
        "wordCount": 500
    },
    {
		"id": "2",
        "title": "My Internship Experience",
        "hook": "A pretty good one",
        "topic": "work",
        "releaseDate": "2023-10-14T00:00:00Z",
        "wordCount": 1000
    }
]
`

	expected := []Post{
		{
			Id:          "1",
			Title:       "How I Made This Site",
			Hook:        "In short Golang, Templ, and RAW DOGGING HTML with the power of ZEUS",
			Topic:       "web dev",
			ReleaseDate: time.Date(2023, time.October, 17, 0, 0, 0, 0, time.UTC),
			WordCount:   500,
		},
		{
			Id:          "2",
			Title:       "My Internship Experience",
			Hook:        "A pretty good one",
			Topic:       "work",
			ReleaseDate: time.Date(2023, time.October, 14, 0, 0, 0, 0, time.UTC),
			WordCount:   1000,
		},
	}

	posts, err := decodePostsJson(strings.NewReader(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	ok := reflect.DeepEqual(expected, posts)
	fmt.Println(expected)
	fmt.Println(posts)
	if !ok {
		t.Log("deep equal failed")
		t.FailNow()
	}
}
