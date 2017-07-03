package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("#searchResultList > li > div > p > a").Each(func(_ int, s *goquery.Selection) {
		//fmt.Println(s)

		url, _ := s.Attr("href")
		fmt.Println(url)
	})
}

func main() {
	url := "http://zozo.jp/men-category/tops/tshirt-cutsew/"
	GetPage(url)
}