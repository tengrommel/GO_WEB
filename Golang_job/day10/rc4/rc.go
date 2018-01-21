package main

import (
	"crypto/rc4"
	"log"
	"io"
	"crypto/md5"
)

func crypto(w io.Writer, r io.Reader, key string)  {
	// 创建cipher
	md5sum := md5.Sum([]byte(key))
	cipher, err := rc4.NewCipher(md5sum[:])
	if err!= nil{
		log.Fatal(err)
	}
	// 创建buf
	buf := make([]byte, 4096)
	for  {
		// 从r里面读取数据到buf
		n, err := r.Read(buf)
		if err == io.EOF{
			break
		}
		// 加密buf
		cipher.XORKeyStream(buf[:n], buf[:n])
		// 把buf写入到w里面
		w.Write(buf[:n])
	}
}

func main() {
	key := "123456"
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil{
		log.Fatal(err)
	}
	buf := []byte("hello")
	cipher.XORKeyStream(buf, buf)
	log.Printf(string(buf))

	{
		cipher, err := rc4.NewCipher([]byte(key))
		if err != nil{
			log.Fatal(err)
		}
		cipher.XORKeyStream(buf, buf)
		log.Printf(string(buf))
	}
}
