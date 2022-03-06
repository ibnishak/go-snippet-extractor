package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	f, err := os.Open("test.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)

	if err != nil {
		log.Fatal(err)
	}
	doc.Find("code.new").Each(func(i int, s *goquery.Selection) {
		file, _ := s.Attr("data-file")
		content := s.Text()
		fmt.Println(file, content)
	})
}
