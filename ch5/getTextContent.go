/*
练习 5.3： 编写函数输出所有text结点的内容。注意不要访问 <script> 和 <style> 元素,因为
这些元素对浏览者是不可见的。
*/
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

/*
 根据 node type，添加 node data 到 text
*/
func getTextContent(text []string, node *html.Node) []string {
	if node.Type == html.TextNode {

		if node.Data != "" {
			//fmt.Println("Content:", node.Data, "ContentLen:", len(node.Data))
			text = append(text, node.Data)
		}
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if n.Data == "script" || n.Data == "style" {
			continue
		}
		text = getTextContent(text, n)
	}

	return text
}

/* Usage: ./fetchUrl https://www.sina.com.cn  | ./getTextContent */
func main() {
	fmt.Println("Get Text Content...")

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("HTML Parse failed")
		os.Exit(1)
	}

	content := getTextContent(nil, doc)
	for _, text := range content {
		fmt.Printf("Text Context:%s\n", text)
	}
}
