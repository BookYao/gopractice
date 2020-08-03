/**
 * @Author: BookYao
 * @Description: anonymous function test. toposort
 * @File:  toposort
 * @Version: 1.0.0
 * @Date: 2020/8/3 23:25
 */

package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string {
	"algorithms":{"data structs"},
	"calculus":{"liner algebra"},
	"compilers":{
		"data structs",
		"formal language",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func toposort(m map[string][]string) [] string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for key:= range keys {
		fmt.Printf("keys[%d]=%s\n", key, keys[key])
	}
	fmt.Println("========================")

	visitAll(keys)

	return order
}

func main() {
	order := toposort(prereqs)
	for key:= range order {
		fmt.Printf("order[%d]=%s\n", key, order[key])
	}
}

  