/**
 * @Author: BookYao
 * @Description:
练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回
另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个
LimitReader函数：
 * @File:  limitReader
 * @Version: 1.0.0
 * @Date: 2020/8/21 11:48
 */

package main

import (
	"fmt"
	"io"
	"os"
)

type LimitReader struct {
	Reader io.Reader
	Limit int
}

func (r *LimitReader) Read(b []byte)(n int, err error) {
	if r.Limit <= 0 {
		return 0, io.EOF
	}

	if len(b) > r.Limit {
		b = b[:r.Limit]
	}

	n, err = r.Reader.Read(b)
	r.Limit -= n
	return
}

func limitReader(r io.Reader, n int) io.Reader {
	return &LimitReader{
		Reader:  r,
		Limit: n,
	}
}

func main() {
	file, err := os.Open("test.txt")
	if  err != nil {
		os.Exit(0)
	}
	defer file.Close()

	lr := limitReader(file, 5)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, buf, string(buf))
}

  