package main

import (
	//"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	//	"net/http"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf(doc)
	doc.Find(".topic").Each(func(i int, s *goquery.Selection) {
		log.Println("hello world")
		content := s.Find(".title a").Text()

		log.Println(content)
		//fmt.Printf("Review %d: %s\n", i, content)
	})
}

func main() {
	ExampleScrape()
}
