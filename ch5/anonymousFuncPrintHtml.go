/**
 * @Author: BookYao
 * @Description: anonymous Function Print Html
 * @File:  anonymousFuncPrintHtml
 * @Version: 1.0.0
 * @Date: 2020/8/7 16:42
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func forEachNode(n *html.Node, prev, post func(node *html.Node)) {
	if prev != nil {
		prev(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, prev, post)
	}

	if post != nil {
		post(n)
	}
}
func anonymousFuncPrintHtml(url string) ([]string, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Get Url: %s-%s", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parse URL: %s. Error:%v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				if strings.Index(a.Val, "java") != -1 {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

func main() {
	if (len(os.Args) != 2) {
		log.Printf("Usage: %s URL string\n", os.Args[0])
		os.Exit(0)
	}

	links , err := anonymousFuncPrintHtml(os.Args[1])
	if err != nil {
		fmt.Printf("anonymousFuncPrintHtml failed. err: %v", err)
	}

	for _, val := range links {
		fmt.Println("val:", val)
	}
}

  