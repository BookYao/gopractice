package main

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/net/html"
)

/*
练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
*/

func visit(links []string, n *html.Node) []string {
	//fmt.Println("Node type:", n.Type, "Node Data:", n.Data)
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "script" || n.Data == "style") {
		//fmt.Println("Node Attr:", n.Attr)
		for _, a :=  range n.Attr {
			//fmt.Println("Key: ", a.Key)
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	/*fmt.Println("fistChild:", n.FirstChild)
	for c := n.FirstChild; c != nil; c =  c.NextSibling {
		links = visit(links, c)
	}*/
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

//Usage: ./fetchUrl https://qq.com | ./getHtmlUrl
func main() {
	fmt.Println("test Get HTML URL")

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("HTML Parse failed!")
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
