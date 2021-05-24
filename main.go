package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	wordlists := []Word{}

	c.OnHTML(".top-g li", func(h *colly.HTMLElement) {
		a := h.DOM.Find("a")
		href, _ := a.Attr("href")

		wt := h.DOM.Find(".pos")
		l := h.DOM.Find(".belong-to")

		wordlists = append(wordlists, Word{
			URL:      href,
			Word:     a.Text(),
			WordType: wt.Text(),
			Level:    l.Text(),
		})
	})

	c.Visit("https://www.oxfordlearnersdictionaries.com/wordlists/oxford3000-5000")

	data, _ := json.Marshal(wordlists)
	err := ioutil.WriteFile("word-list.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

type Word struct {
	URL      string
	Word     string
	WordType string
	Level    string
}
