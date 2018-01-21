package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8021")
	if err!=nil{
		log.Fatal(err)
	}

	conn.Write([]byte("GET /home/teng/java_error_in_GOLAND_2078.log"))
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(buf[:n])
}
