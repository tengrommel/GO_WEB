package main

import (
	"os"
	"net/url"
	"log"
	"fmt"
)

func main() {
	s := os.Args[1]
	u, err := url.Parse(s)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("scheme:",u.Scheme)
	fmt.Println("host", u.Host)
	fmt.Println("path", u.Path)
	fmt.Println("queryString", u.RawQuery)
	fmt.Println("user", u.User)
	fmt.Println("xx", u.Fragment)
}
