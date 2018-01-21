package main

import (
	"time"
	"fmt"
	"strconv"
)

func Benchmark(start time.Time)  {
	duration := time.Now().Nanosecond()/1e3 - start.Nanosecond()/1e3
	fmt.Println("Duration: " + strconv.Itoa(duration) + "ms")
}
