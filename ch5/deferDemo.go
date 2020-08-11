/**
 * @Author: BookYao
 * @Description: defer demo, calculate exec time of function
 * @File:  deferDemo
 * @Version: 1.0.0
 * @Date: 2020/8/11 23:10
 */

package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)

	return func() {
		log.Printf("exit %s(%s)\n", msg, time.Since(start))
	}
}

func slowOperation() {
	defer trace("slowOperation")()

	time.Sleep(5 * time.Second)
}

func main() {
	slowOperation()
}

  