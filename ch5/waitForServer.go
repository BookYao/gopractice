/**
 * @Author: BookYao
 * @Description: 主要测试程序异常时的错误处理
 * @File:  waitForServer
 * @Version: 1.0.0
 * @Date: 2020/7/26 22:19
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	//fmt.Println("timeout:", timeout, "deadline:", deadline)

	for tries := 0; time.Now().Before(deadline) ;tries ++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.Printf("Server not responding. err:%s. retrying...", err)
		time.Sleep(time.Second << tries)
	}

	return fmt.Errorf("Server URL(%s) failed to respond after %s", url, timeout)
}

func main() {
	log.Printf("Wait for Server response...")

	if len(os.Args) != 2 {
		log.Printf("Usage: %s URL!", os.Args[0])
		os.Exit(0)
	}

	fmt.Println(waitForServer(os.Args[1]))
}

  