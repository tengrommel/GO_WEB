package main

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"errors"
	"strings"
)

func fetch(url string) ([]string, error) {
	//url := "https://www.douban.com/"
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil{
		return nil, err
	}
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		link, ok := selection.Attr("src")
		if ok && strings.HasSuffix(link, ".jpg") {
				urls = append(urls, link)
		}
	})
	return urls, nil
}

func cleanUrls(u string, urls []string) []string {
	if u ==""{
		urls = append(urls, u)
	}
	return urls
}



func main()  {
	url := "https://www.douban.com/"
	urls, err := fetch(url)
	if err!=nil{
		log.Fatal(err)
	}
	for _, u := range urls{
		fmt.Println(u)
	}
}