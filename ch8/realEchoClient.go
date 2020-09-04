/**
 * @Author: BookYao
 * @Description:
 * @File:  readEchoClient
 * @Version: 1.0.0
 * @Date: 2020/9/4 16:03
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
		log.Printf("IO Copy Failed. err:%v\n", err)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s addr:port\n", os.Args[0])
		os.Exit(0)
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Printf("Connect Failed. err:%v\n", err)
		os.Exit(0)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

  