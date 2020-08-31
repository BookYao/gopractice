/**
 * @Author: BookYao
 * @Description: 根据错误号返回错误串的demo
 * @File:  errorString
 * @Version: 1.0.0
 * @Date: 2020/8/31 11:28
 */

package main

import (
	"fmt"
	"log"
)

type Errno uintptr

var errorString = [...]string {
	1: "error string 1",
	2: "error string 2",
	3: "error string 3",
	4: "error string 4",
}

func (e Errno) errString() string {
	if int(e) > 0 && int(e) < len(errorString) {
		return errorString[e]
	}

	return fmt.Sprintf("Not found errno: %d", e)
}

func (e Errno) errString1(e1 Errno) string {
	if int(e1) > 0 && int(e1) < len(errorString) {
		return errorString[e1]
	}

	return fmt.Sprintf("Not found errno: %d", e1)
}

func main() {
	var err Errno = 1
	log.Printf("Get error1: %s\n", err.errString())

	err = 2
	log.Printf("Get error1: %s\n", err.errString())

	err = 10
	log.Printf("Get error1: %s\n", err.errString())

	err = 19
	log.Printf("Get error1: %s\n", err.errString1(err))

	err = 4
	log.Printf("Get error1: %s\n", err.errString1(err))

}

  