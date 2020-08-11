/**
 * @Author: BookYao
 * @Description: 练习5.16：编写多参数版本的strings.Join
 * @File:  argsStringJoin
 * @Version: 1.0.0
 * @Date: 2020/8/11 11:15
 */

package main

import (
	"fmt"
	"strings"
)

func stringJoin(args []string, delim string) string {
	length := len(args)
	if length == 0 {
		return ""
	}

	var str string
	for i, val := range args {
		if (i == length -1 ) {
			str += val
		} else {
			str += val + delim
		}
	}

	return str
}

func main() {
	s1 :=[]string{"first", "second", "third"}
	fmt.Println("join string:", stringJoin(s1, ":"))
	s := []string{"first", "second", "third"}
	fmt.Println("stringjoin:", strings.Join(s, ":"))
}

  