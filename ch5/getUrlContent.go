/**
 * @Author: BookYao
 * @Description: 保存访问的html页面内容
 * @File:  getUrlContent
 * @Version: 1.0.0
 * @Date: 2020/8/11 23:40
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (str string, len int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Http Get Failed. url:%s\n", url)
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	//fmt.Println("local:", local)
	if local == "/" || local == "." {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err := io.Copy(f, resp.Body)
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}

	return local, n, err
}

/*Usage: ./getUrlContent https://www.sina.com
or: ./getUrlContent https://www.baidu.com */
func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s url", os.Args[0])
		os.Exit(0)
	}
	filename, len, err:= fetch(os.Args[1])
	if err != nil {
		log.Println("fetch failed. strerror:", fmt.Errorf("%v", err))
	}
	log.Printf("File(%s) length:%d\n", filename, len)
}

  