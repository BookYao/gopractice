/**
 * @Author: BookYao
 * @Description:
 * @File:  links
 * @Version: 1.0.0
 * @Date: 2020/9/14 22:39
 */

package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parse %s As html. error:%v\n", url, err)
	}

	var links []string
	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key != "herf" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				links = append(links, link)
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}


func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for elem := n.FirstChild; elem != nil; elem = elem.NextSibling {
		forEachNode(elem, pre, post)
	}

	if post != nil {
		post(n)
	}
}

  