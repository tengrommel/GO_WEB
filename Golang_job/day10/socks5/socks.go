package main

import (
	"net"
	"flag"
	"log"
	"bufio"
	"errors"
	"io"
	"encoding/binary"
	"fmt"
)

// 1.握手
// 2.获取客户端代理的请求
// ３．开始代理

func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("version: %d", version)
	if version != 5{
		return "", errors.New("bad version")
	}
	cmd, _ := r.ReadByte()
	log.Printf("cmd:%d", cmd)
	if cmd != 1{
		return "", errors.New("bad cmd")
	}
	r.ReadByte()
	addrtype, _ := r.ReadByte()
	log.Printf("addr type:%d", addrtype)
	if addrtype != 3{
		return "", errors.New("bad addr type")
	}

	// 读取一个字节的数据，代表后面紧跟着的域名的长度
	// 读取n个字节得到域名，n根据上一步得到的结果来决定
	// addrlen
	// addr
	addrlen, _ := r.ReadByte()
	addr := make([]byte, addrlen)
	io.ReadFull(r, addr)
	log.Printf("addr: %s", addr)

	var port int16
	binary.Read(r, binary.BigEndian, &port)

	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handshake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Printf("verson: %d", version)
	if version != 5{
		return errors.New("bad version")
	}
	nmethods, _ := r.ReadByte()
	log.Printf("n methods: %d", nmethods)

	buf := make([]byte, nmethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)

	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handshake(r, conn)
	addr, _ := readAddr(r)
	log.Printf("addr:%s", addr)
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)

	// 开始代理

}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", ":8021")
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
