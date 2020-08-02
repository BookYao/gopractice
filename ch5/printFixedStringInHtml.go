/**
 * @Author: BookYao
 * @Description:
练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止
forEachNoded的遍历。使用修改后的代码编写ElementByID函数，根据用户输入的id查找第
一个拥有该id元素的HTML元素，查找成功后，停止遍历。
func ElementByID(doc *html.Node, id string) *html.Node

 * @File:  printFixedStringInHtml
 * @Version: 1.0.0
 * @Date: 2020/8/2 9:36
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

type callback func(*html.Node, string) bool

func startElem(n *html.Node, str string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			//fmt.Println("a.kay:", a.Key, "str:", str)
			if a.Key == str {
				fmt.Printf("Start Fixed String Value: %s\n", a.Val)
				return true
			}
		}
	}

	return false
}

func ElementByID(node *html.Node, idstr string, pre callback) *html.Node {
	if pre != nil {
		if pre(node, idstr) {
			return node
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(c, idstr, pre)
	}

	return node
}

func parseHtml(url string) (*html.Node, error) {
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

	return doc, nil
}


/*
Usage: ./printFixedStringInHtml https://qq.com href
or ./printFixedStringInHtml https://qq.com src
or ./printFixedStringInHtml https://qq.com class
*/
func main() {
	log.Println("Print HTML Node.")

	if (len(os.Args) != 3) {
		log.Printf("Usage: %s URL string\n", os.Args[0])
		os.Exit(0)
	}

	doc, err := parseHtml(os.Args[1])
	if err != nil {
		log.Printf("Print Html Failed: url:%s-err:%v\n", os.Args[1], err)
	}

	ElementByID(doc, os.Args[2], startElem)
}

  