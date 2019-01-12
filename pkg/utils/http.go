package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
	"github.com/matryer/try"
)

func GetDocument(url string) *goquery.Document {
	doc, _ := goquery.NewDocument("")
	err := try.Do(func(attempt int) (retry bool, err error) {
		retry = attempt < 5 // try 5 times
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(fmt.Sprintf("panic: %v", r))
			}
		}()
		doc, err = request(url)
		return
	})
	if err != nil {
		log.Fatalln("error:", err)
	}

	return doc
}

func request(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	utfBody, err := iconv.NewReader(res.Body, "ISO-8859-9", "utf-8")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
