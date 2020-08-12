/**
 * @Author: BookYao
 * @Description: 增加 panic demo
 * @File:  panicDemo
 * @Version: 1.0.0
 * @Date: 2020/8/12 16:17
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, nil)
	}

	if  post != nil {
		post(n)
	}
}

func soleTile(node *html.Node)(title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("Multil element")
		default:
			panic(p)
		}
	}()

	forEachNode(node, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println("title:", title)
			if title != "" {
				fmt.Println("Titl Has value. Panic happen...")
				panic(bailout{})
			}

			title = n.FirstChild.Data
			fmt.Println("firstData, title:", title)
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {
	if (len(os.Args) != 2) {
		log.Printf("Usage: %s URL string\n", os.Args[0])
		os.Exit(0)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Printf("Http Get Failed. url:%s", os.Args[1])
		os.Exit(0)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(fmt.Errorf("Get Url: %s-%s", os.Args[1], resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Parse URL: %s. Error:%v", os.Args[1], err))
	}

	title, err := soleTile(doc)
	if err != nil {
		log.Printf("SoleTitle Failed. strerror:%v\n", err)
	}
	log.Println("title:", title)
}
  