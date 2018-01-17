package main

import (
	"fmt"
	"path"
)

func main() {
	s := "segmentfault.com/q/1010000006670932"
	dir := path.Dir(s)
	name := path.Base(s)
	fullname := path.Join(dir, name)
	fmt.Println(fullname)
}
