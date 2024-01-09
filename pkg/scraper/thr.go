package scraper

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

type Post struct {
	Title    string
	Url      string
	ImageUrl string
	Content  string
	Intro    string
}

func HollywoodReporterPosts() []Post {
	var posts []Post
	var i = 0

	c := colly.NewCollector()
	url := "https://www.hollywoodreporter.com/c/movies/movie-news/"

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("div.story", func(e *colly.HTMLElement) {
		if i < 3 {
			time.Sleep(1 * time.Second)
			var post Post
			post.Title = e.ChildText("a.c-title__link")
			post.ImageUrl = e.ChildAttr("img.c-lazy-image__img", "data-lazy-src")
			post.Url = e.ChildAttr("a.c-title__link", "href")
			post.Intro, post.Content = hollywoodReporterContent(post.Url)
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

func hollywoodReporterContent(url string) (intro string, content string) {

	time.Sleep(1 * time.Second)

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("main", func(e *colly.HTMLElement) {
		intro = e.ChildText("p.article-excerpt")
		content = e.ChildText("p.paragraph")
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return
}
