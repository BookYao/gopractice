/**
 * @Author: BookYao
 * @Description: 练习 5.5： 实现countWordsAndImages。
 * @File:  countWordsAndImage
 * @Version: 1.0.0
 * @Date: 2020/7/25 22:37
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

func urlParse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Parse URL Failed");
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err := fmt.Errorf("Parse HTML Failed: %s", err)
		return nil, err
	}

	return doc, nil
}

func visit(texts []string, image int, n *html.Node) ([]string, int) {
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		image++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		texts, image = visit(texts, image, c)
	}

	return texts, image
}

func countWordAndImageFunc(node *html.Node) (words, images int) {
	texts, images := visit(nil, 0, node)

	for _, v := range texts {
		v = strings.Trim(strings.TrimSpace(v), "\r\n")
		if v == "" {
			continue
		}
		words += strings.Count(v, "")
	}

	return
}

func countWordsAndImage(url string) (words, image int) {
	doc, err := urlParse(url)
	if err != nil {
		fmt.Println("Parse URL failed")
		return
	}

	words, image = countWordAndImageFunc(doc)

	return
}

/*
Usage: ./countWordsAndImage  http://sina.com.cn
*/
func main() {
	if (len(os.Args) != 2) {
		log.Printf("Usage: %s URL", os.Args[0])
		os.Exit(0)
	}

	words, image := countWordsAndImage(os.Args[1])
	fmt.Printf("words:%d-image:%d\n", words, image)
}


  