/**
 * @Author: BookYao
 * @Description: 练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
 * @File:  replaceString
 * @Version: 1.0.0
 * @Date: 2020/8/2 8:53
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

type callback func(s string) string

func dstStr(str string) string {
	return strings.ToUpper(str)
}

func replaceString(s string, f callback) {
	str := f("foo")
	s = strings.Replace(s, "foo", str, -1)
	fmt.Println("Finish String Replace. String:", s)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s String\n", os.Args[0])
		os.Exit(0)
	}
	replaceString(os.Args[1], dstStr)
}

  