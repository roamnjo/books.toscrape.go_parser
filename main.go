package main

import (
	"net/http"
	"time"

	"github.com/roamnjo/books.toscrape.go_parser/web_parser"
)

func main() {
	URL := "https://books.toscrape.com/catalogue/"
	currentPage := "page-1.html"

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	web_parser.Parse(URL, currentPage, client)
}
