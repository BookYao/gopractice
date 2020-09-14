/**
 * @Author: BookYao
 * @Description:
 * @File:  crawl
 * @Version: 1.0.0
 * @Date: 2020/9/14 23:11
 */

package main

import (
	"fmt"
	links2 "gopl.io/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string{
	fmt.Println(url)
	tokens <- struct {}{}
	links, err := links2.Extract(url)
	<- tokens
	if err != nil {
		log.Printf("link extrace failed. url:%s-err:%v\n", url, err)
		return nil
	}

	return links
}

func main() {
	worklist := make(chan []string)

	var n int

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <- worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(url string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

  