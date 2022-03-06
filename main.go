package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	dat, err := ioutil.ReadFile("test.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	var buf bytes.Buffer
	if err := md.Convert(dat, &buf); err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(buf.String()))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("code.new").Each(func(i int, s *goquery.Selection) {
		file, _ := s.Attr("data-file")
		content := s.Text()
		fmt.Println(file, content)
	})

}
