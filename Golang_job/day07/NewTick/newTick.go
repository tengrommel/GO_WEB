package main

import (
	"time"
	"fmt"
)

func main() {
	defer func() { // 运行时的错误无法捕获
		recover()
	}()

	timer := time.NewTicker(time.Second)
	cnt := 0
	for _ = range timer.C {
		cnt++
		if cnt > 10{
			timer.Stop()
			// Stop不会关闭所以必须手动退出
			return
		}
		fmt.Println("hello")
	}
}
