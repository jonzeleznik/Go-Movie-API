package scraper

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

type Post struct {
	Title    string
	ImageUrl string
}

func HollywoodReporter() []Post {
	var posts []Post
	var i = 0

	c := colly.NewCollector()
	url := "https://www.hollywoodreporter.com/c/movies/movie-news/"

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("div.story", func(e *colly.HTMLElement) {
		if i < 2 {
			time.Sleep(1 * time.Second)
			var post Post
			post.Title = e.ChildText("h3")
			post.ImageUrl = e.ChildAttr("img.c-lazy-image__img", "data-lazy-src")
			posts = append(posts, post)
			i++
		}
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return posts
}
