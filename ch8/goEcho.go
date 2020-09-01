/**
 * @Author: BookYao
 * @Description:
 * @File:  goruntineEcho
 * @Version: 1.0.0
 * @Date: 2020/9/1 23:20
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("2006/1/2 03:04:05\n"))
		if err != nil {
			log.Println("Write string failed.")
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed. %v", err)
			continue
		}

		go handleConn(conn)
	}
}

  