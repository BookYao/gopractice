/**
 * @Author: BookYao
 * @Description:
 * @File:  channelDemo
 * @Version: 1.0.0
 * @Date: 2020/9/7 15:33
 */

package main

import (
	"fmt"
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
	conn, err := net.Dial("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Printf("Connect Failed. err:%v", err)
		os.Exit(0)
	}

	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Printf("Read Failed. err:%v\n", err)
		}
		fmt.Println("Done...")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	fmt.Println("main gorounte finish...")
	<- done
}

  