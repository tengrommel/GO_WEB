package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	md5sum := md5.Sum([]byte("hello"))
	fmt.Printf("%x\n", md5sum)

	md5sum1 := md5.Sum([]byte("hello1"))
	if md5sum == md5sum1{
		fmt.Println("error")
	}
}
