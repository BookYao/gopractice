/**
 * @Author: BookYao
 * @Description: anonymous Function Print Html
 * @File:  anonymousFuncPrintHtml
 * @Version: 1.0.0
 * @Date: 2020/8/7 16:42
 * update1: 增加 htmlSave
 练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。只
保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页
面
 */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func breadFist(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func htmlSave(u string) bool {
	urlObj, _ := url.Parse(u)
	urlPath := "/var/tmp/crawl/" + urlObj.Host

	//fmt.Println("urlObj.path:", urlObj.Path)
	if urlObj.Path == "" || urlObj.Path == "/" {
		urlObj.Path = "/index.html"
	}

	//fmt.Println("urlpath:", urlPath)
	fileinfo, err := os.Stat(urlPath)
	if err != nil {
		os.MkdirAll(urlPath, 0755)
	}
	fileinfo = fileinfo

	fileName := urlPath + urlObj.Path
	fmt.Println("FileName:", fileName)

	fp, err := os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("openfile failed, filename:%s", fileName)
		return false
	}

	resp, err := http.Get(u)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("Get Http resp failed. err:%v", err)
		return false
	}

	body, _ := ioutil.ReadAll(resp.Body)

	io.WriteString(fp, string(body))
	resp.Body.Close()
	fp.Close()

	body = nil
	return true
}

func crawl(url string) []string {
	//fmt.Println("URL:", url)
	go htmlSave(url)
	lists, err := anonymousFuncPrintHtml(url)
	if err != nil {
		fmt.Printf("anonymousFuncPrintHtml failed. err: %v", err)
	}

	return lists
}

/*
Usage: ./anonymousFuncPrintHtml  https://www.sina.com
*/
func main() {
	if (len(os.Args) != 2) {
		log.Printf("Usage: %s URL string\n", os.Args[0])
		os.Exit(0)
	}

	/*links , err := anonymousFuncPrintHtml(os.Args[1])
	if err != nil {
		fmt.Printf("anonymousFuncPrintHtml failed. err: %v", err)
	}

	for _, val := range links {
		fmt.Println("val:", val)
	}*/

	breadFist(crawl, os.Args[1:])
}

  