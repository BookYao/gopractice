package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func outline(url []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		url = append(url, node.Data)
		//fmt.Println("URL:", url)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		url = outline(url, n)
	}
	return url
}

func calSameStringCount(str []string) map[string]int {
	result := make(map[string]int)

	for _, val := range str {
		result[val]++
	}

	return result
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("HTML Parse failed")
		os.Exit(1)
	}

	url := outline(nil, doc)
	fmt.Println("url:", url)

	fmt.Println("===========================")
	result := calSameStringCount(url)
	for str, count := range result {
		fmt.Printf("str:%s-count:%d\n", str, count)
	}
}
