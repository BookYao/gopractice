/* * * * * * * * * * * * * * * * * * * * * * * * * * *
* 练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。
* * * * * * * * * * * * * * * * * * * * * * * * * * */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

/*
 HTML doc 先存入到 slice， 然后再统计 elem 个数
 */
func outline(elem []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		elem = append(elem, node.Data)
		//fmt.Println("URL:", url)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		elem = outline(elem, n)
	}
	return elem
}

func calSameStringCount(str []string) map[string]int {
	result := make(map[string]int)

	for _, val := range str {
		result[val]++
	}

	return result
}

/*
递归方法统计 elem 个数
*/
func calSameStringCount2(result map[string] int, node *html.Node) map[string]int {
	if node.Type == html.ElementNode {
		result[node.Data]++
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		result = calSameStringCount2(result, n)
	}

	return result
}

/* Usage: ./fetchUrl https://qq.com | ./outline */
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("HTML Parse failed")
		os.Exit(1)
	}

	//function1, Save all ElementNode in elem, and calculate elem count
	/*elem := outline(nil, doc)
	result := calSameStringCount(elem)
	for str, count := range result {
		fmt.Printf("string:%s---count:%d\n", str, count)
	}*/

	//function2, indirect calculate elem count by HTML doc.
	result := make(map[string] int)
	calSameStringCount2(result, doc)
	for str, count := range result {
		fmt.Printf("string:%s---count:%d\n", str, count)
	}
}
