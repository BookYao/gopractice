package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func visit(link []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a :=  range node.Attr {
			//fmt.Println("Key: ", a.Key)
			if a.Key == "href" {
				link = append(link, a.Val)
			}
		}
	}

	if node.FirstChild != nil {
		link = visit(link, node.FirstChild)
	}

	if node.NextSibling != nil {
		link = visit(link, node.NextSibling)
	}

	return link
}

func findLinks(url string) ([]string, error) {
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

	return visit(nil, doc), nil
}

/*
Get HTML URL Version 2
Usage: ./a.out  URL1  URL2
e.g.: ./getHtmlUrlVer2 http://sina.com.cn http://qq.com
*/
func main() {
	fmt.Println("Get Html Url, version 2")
	if len(os.Args) < 1 {
		fmt.Printf("Usage: %s URL", os.Args[0])
		os.Exit(0)
	}

	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "findLinks error: %v", err)
			continue
		}

		for _, link := range links {
			fmt.Println("Link: ", link)
		}
	}
}
