/**
 * @Author: BookYao
 * @Description: Print HTML Node
 * @File:  printHtml
 * @Version: 1.0.0
 * @Date: 2020/7/30 23:40
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

var depth int

func preElem(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
		depth++
	}
}

func postElem(n *html.Node) {
	if (n.Type == html.ElementNode) {
		depth--
		fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
	}
}

func printHtmlContent(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		printHtmlContent(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

func printHtml(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("Get Url: %s-%s", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("Parse URL: %s. Error:%v", url, err)
	}

	printHtmlContent(doc, preElem, postElem)

	return nil
}


/*
Usage: printHtml  https://www.baidu.com
or  pritnHtml https://qq.com
*/
func main() {
	log.Println("Print HTML Node.")

	if (len(os.Args) != 2) {
		log.Printf("Usage: %s URL.\n", os.Args[0])
		os.Exit(0)
	}

	err := printHtml(os.Args[1])
	if err != nil {
		log.Printf("Print Html Failed: url:%s-err:%v\n", os.Args[1], err)
	}
}



  