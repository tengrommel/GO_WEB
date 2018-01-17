package main

import "fmt"
/*
通过协程和管道来进行加速程序的执行
 */
func sum(s []int, c chan int)  {
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
