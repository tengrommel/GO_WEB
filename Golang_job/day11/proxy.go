package day11

import (
	"net"
	"flag"
	"io"
	"log"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com", "f")
)

func handleConn(conn net.Conn)  {
	// 建立到目标服务器的连接
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", *target)
	if err!=nil{
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go　接收客户端的数据发送到remote，直到conn的eof，关闭remote
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	// go 接收remote的数据，发送到客户端，直到remote的eof，关闭conn
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	// 等待两个协程结束
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}
