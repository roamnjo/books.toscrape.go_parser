package web_parser

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Book struct {
	Title  string
	Rating string
	Price  string
	Link   string
}

func Parse(url, currentPage string, client *http.Client) {
	for {
		fullURL := url + currentPage

		doc, err := FetchPage(fullURL, client)
		if err != nil {
			log.Fatal("Error:", err)
		}

		books := ParseBooks(doc)
		CheckBookRating(books)

		nextPage := FindNext(doc)
		if nextPage == "" {
			break
		}
		currentPage = nextPage
	}
}

func FetchPage(url string, client *http.Client) (*goquery.Document, error) {
	res, err := client.Get(url)
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
		return nil, nil
	}
	return doc, nil
}

func ParseBooks(doc *goquery.Document) []Book {
	var books []Book

	doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		b := Book{}

		b.Title, _ = s.Find("h3 a").Attr("title")
		b.Rating, _ = s.Find("p.star-rating").Attr("class")
		b.Price = s.Find("p.price_color").Text()
		link, _ := s.Find("h3 a").Attr("href")
		b.Link = "https://books.toscrape.com/" + strings.TrimPrefix(link, "../")

		books = append(books, b)
	})
	return books
}

func CheckBookRating(books []Book) {
	for _, book := range books {
		if strings.Contains(book.Rating, "Four") || strings.Contains(book.Rating, "Five") {
			fmt.Println(book)
		}
	}
}

func FindNext(doc *goquery.Document) string {
	next := ""

	doc.Find("li.next a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			next = href
		}
	})
	return next
}
