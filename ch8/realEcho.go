/**
 * @Author: BookYao
 * @Description:
 * @File:  realEcho
 * @Version: 1.0.0
 * @Date: 2020/9/4 15:45
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func realEcho(conn net.Conn, msg string, delay time.Duration) {
	fmt.Fprintf(conn, "%s\n", strings.ToUpper(msg))
	time.Sleep(delay)
	fmt.Fprintf(conn, "%s\n", msg)
	time.Sleep(delay)
	fmt.Fprintf(conn, "%s\n", strings.ToLower(msg))
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go realEcho(conn, input.Text(), 1 * time.Second)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s serveraddr:port\n", os.Args[0])
		os.Exit(0)
	}

	listener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Printf("Listen Failed. err:%v\n", err)
		os.Exit(0)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept Failed. err:%v\n", err)
			continue
		}

		go handleConn(conn)
	}
}

  