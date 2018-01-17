package main

import (
	"time"
	"fmt"
)

/*
利用多线程进行排序
通过channel进行协程之间的通信
 */

func main() {
	s := []int{2,7,1,6,4}
	for _, n := range s{
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(n)
	}
	time.Sleep(10 * time.Second)
}
