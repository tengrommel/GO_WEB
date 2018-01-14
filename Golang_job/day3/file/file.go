package main

import (
	"os"
	"log"
	"fmt"
	//"bufio"
	//"io"
)

func main() {
	// os.Create相当于 os.
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err!=nil{
		log.Fatal(err)
	}
	//f.WriteString("hello\n")
	//f.Seek(0, os.SEEK_SET)
	//f.WriteString("ss")
	////buf := make([]byte, 1024)
	////f.Read(buf)
	//r := bufio.NewReader(f)
	//for {
	//	line, err := r.ReadString('\n')
	//	if err == io.EOF{
	//		break
	//	}
	//	fmt.Print(line)
	//}
	f.Seek(3, os.SEEK_SET)
	buf := make([]byte, 2)
	f.Read(buf)
	fmt.Print(buf)
	f.Close()
}
