package main

import (
	"fmt"
)
/*
可以用两个chan改变
*/
func joint(s []string, jointResult chan string)  {
	a := ""
	for _, r :=range s{
		a += r
	}
	jointResult <- a
}

func main() {
	backChan := make(chan string)
	s := []string{"hello", "golang", "c++", "world"}
	go joint(s[:len(s)/2], backChan)
	go joint(s[len(s)/2:], backChan)
	x, y := <-backChan, <-backChan
	fmt.Println(x+y)
}
