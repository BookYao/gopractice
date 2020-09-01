/**
 * @Author: BookYao
 * @Description:
 * @File:  goEchoClient
 * @Version: 1.0.0
 * @Date: 2020/9/1 23:43
 */

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(w io.Writer, r io.Reader) {
	_, err := io.Copy(w, r)
	if err != nil {
		log.Println("io copy failed. %v", err)
	}
}

func main() {
	conn, err := net.Dial("tcp","localhost:8003")
	if err != nil {
		log.Fatal("echo client connect failed")
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

  