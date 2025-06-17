package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://books.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("error: StatusCode:", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("error getting document:", err)
		return
	}

	doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find("h3 a").Attr("title")
		rating, _ := s.Find("p.star-rating").Attr("class")
		price := s.Find("p.price_color").Text()
		html_link, _ := s.Find("h3 a").Attr("href")
		link := "https://books.toscrape.com/" + html_link

		fmt.Printf("title: %s\n%s\nprice: %s\n%s\n\n", title, rating, price, link)
	})
}
