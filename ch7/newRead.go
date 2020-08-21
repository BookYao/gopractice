/**
 * @Author: BookYao
 * @Description:
练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型
的值（和其它值） 。实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的
HTML解析器（§5.2）
 * @File:  newRead
 * @Version: 1.0.0
 * @Date: 2020/8/20 23:17
 */

package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	data string
	current int
}

func (sr *StringReader) Read(b []byte)(n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}

	// copy() guarantee copy min(len(b),len(sr.data[sr.current:])) bytes
	n = copy(b, sr.data[sr.current:])
	if sr.current += n; sr.current >= len(sr.data) { // 已读完
		err = io.EOF
	}

	return
}

func NewReader(str string) *StringReader {
	sr := new(StringReader)
	sr.data = str
	return sr
}

func main() {
	str := "hello world"
	sr := NewReader(str)

	data := make([]byte, 10)
	n, err := sr.Read(data[:0])
	for err == nil {
		n, err = sr.Read(data)
		fmt.Println(n, string(data[0:n]))
	}

    fmt.Println("....")
}

  
