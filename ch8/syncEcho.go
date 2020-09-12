/**
 * @Author: BookYao
 * @Description:
练习 8.4： 修改reverb2服务器，在每一个连接中使用sync.WaitGroup来计数活跃的echo
goroutine。当计数减为零时，关闭TCP连接的写入，像练习8.3中一样。验证一下你的修改版
netcat3客户端会一直等待所有的并发“喊叫”完成，即使是在标准输入流已经关闭的情况下。
 * @File:  syncEcho
 * @Version: 1.0.0
 * @Date: 2020/9/12 22:55
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func handleConn(conn net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(conn)
	for input.Scan() {
		wg.Add(1)
		go func(conn net.Conn, msg string, delay time.Duration) {
			defer wg.Done()
			fmt.Fprintf(conn, "%s\n", strings.ToUpper(msg))
			time.Sleep(delay)
			fmt.Fprintf(conn, "%s\n", msg)
			time.Sleep(delay)
			fmt.Fprintf(conn, "%s\n", strings.ToLower(msg))
		} (conn, input.Text(), 1*time.Second)
	}

	wg.Wait()
	conn.Close()
}

/*Usage: ./syncEcho 127.0.0.1:8003*/
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



  