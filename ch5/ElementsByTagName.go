/**
 * @Author: BookYao
 * @Description: 练习5.17：编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数
量的标签名，返回与这些标签名匹配的所有元素
 * @File:  ElementsByTagName
 * @Version: 1.0.0
 * @Date: 2020/8/11 12:13
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

var nodes []*html.Node

func elementsByTagName(node *html.Node, args ...string) []*html.Node{
	newNode := node
	for _, arg := range args {
		if node.Type == html.ElementNode && node.Data == arg {
			fmt.Println("111111: node.type:", node.Type, "node.Data:", node.Data, "arg:", arg)
			nodes = append(nodes, newNode)
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		elementsByTagName(c, args...)
	}

	return nodes
}

func imgHtmlNode(n *html.Node) {
	//fmt.Println("n.Data:", n.Data)
	for _, a := range n.Attr {
		//fmt.Printf("a.key:%s---a.Val:%s\n", a.Key, a.Val)
		if a.Key == "src" {
			fmt.Printf("Img:%s\n", a.Val)
		}
	}
}

func headHtmlNode(n *html.Node) {
	fmt.Println("n.Data:", n.Data)
}

func forEachNode(n []*html.Node, f func(node *html.Node)) {
	for index, _ := range n {
		f(n[index])
	}
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

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println(fmt.Errorf("Get Url: %s-%s", os.Args[1], resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(fmt.Errorf("Parse URL: %s. Error:%v", os.Args[1], err))
	}

	fmt.Println("11111")

	//./ElementsByTagName https://www.sina.com
	elementsByTagName(doc, "img")
	elementsByTagName(doc, "h1")

	forEachNode(nodes, imgHtmlNode)
	fmt.Println("===========================")
	forEachNode(nodes, headHtmlNode)
}