package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

const (
	SELECTOR   = "code.snippet" //Tagname = code, Class = snippet
	DEFAULTDIR = "."            //Directory to extract snippet to.
)

func main() {

	var file, targetDir string
	flag.StringVar(&file, "f", "", "File from which snippets should be extracted.")
	flag.StringVar(&targetDir, "d", DEFAULTDIR, "Directory to which snippets should be extracted")
	flag.Parse()

	//File is a must
	if file == "" {
		log.Fatal("No file given")
	}

	data, err := openFile(file)
	if err != nil {
		log.Fatal("Error in opening file", err.Error())
	}

	data.Find(SELECTOR).Each(func(i int, s *goquery.Selection) {
		filename, _ := s.Attr("data-file")
		targetFile := filepath.Join(targetDir, filename)
		content := s.Text()

		f, err := os.Create(targetFile)
		if err != nil {
			log.Fatal("Error in creating file", err.Error())
		}
		defer f.Close()
		_, err = f.WriteString(content)
		if err != nil {
			log.Fatal("Error in writing file", err.Error())
		}
	})
}

func openFile(file string) (*goquery.Document, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, err
	}
	return data, err
}
