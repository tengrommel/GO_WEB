package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
)

type Message struct {
	Content string `json:"content"`
	FurtherContent string `json:"further_content"`
}

var url string = "http://localhost:8080"


func Request(url string, target func(string) (*http.Response, error)) Message {
	res, err := target(url)
	if err != nil{
		log.Fatal("Didn't work:" + err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	m := Message{}
	json.Unmarshal(body, &m)
	return m
}

func main() {
	Request(url, http.DefaultClient.Get)
}
