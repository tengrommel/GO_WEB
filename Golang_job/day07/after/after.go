package main

import (
	"time"
	"fmt"
)

func main() {
	c := time.After(time.Second * 3)
	<-c
	fmt.Println("done")
}
