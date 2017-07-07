package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("#searchResultList > li > div > p > a").Each(func(_ int, s *goquery.Selection) {
		//fmt.Println(s)

		path, _ := s.Attr("href")
		url = "http://zozo.jp" + path

		fmt.Println(url)

		page, _ := goquery.NewDocument(url)

		title, _ := page.Find("title").Html()

		text, _ := shiftJIS2UTF8(title)

		fmt.Println(text)
	})
}

func shiftJIS2UTF8(str string) (string, error) {
	strReader := strings.NewReader(str)
	decodedReader := transform.NewReader(strReader, japanese.ShiftJIS.NewDecoder())
	decoded, err := ioutil.ReadAll(decodedReader)
	if err != nil {
		return "", err
	}
	return string(decoded), err
}

func main() {
	url := "http://zozo.jp/men-category/tops/tshirt-cutsew/"
	GetPage(url)
}