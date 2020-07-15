package main

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	//fmt.Println("Node type:", n.Type, "Node Data:", n.Data)
	if n.Type == html.ElementNode && n.Data == "a" {
		//fmt.Println("Node Attr:", n.Attr)
		for _, a :=  range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	//fmt.Println("fistChild:", n.FirstChild)
	//for c := n.FirstChild; c != nil; c =  c.NextSibling {
	//	links = visit(links, c)
	//}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
func main() {
	fmt.Println("test Get HTML URL")

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("HTML Parse failed!")
	}

	fmt.Println("Start visit...")
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
