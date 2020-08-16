/**
 * @Author: BookYao
 * @Description:
 * @File:  listSum
 * @Version: 1.0.0
 * @Date: 2020/8/16 18:05
 */

package main

import "fmt"

type object interface{}

type Node struct {
	val object
	next *Node
}

type List struct  {
	listnode *Node
}

func (list *List) Add(data object) {
	node := &Node{data, nil}
	node.next = list.listnode
	list.listnode = node
}

func (list *List) Sum() int {
	sum := 0
	for node := list.listnode; node != nil; node = node.next {
		switch v := node.val.(type) {
		case int:
			sum += v
		default:
			continue
		}
	}
	return sum
}

func (list *List) travel(f func(n *Node)) {
	for node := list.listnode; node != nil; node = node.next {
		f(node)
	}
}

func travelInt(n *Node) {
	fmt.Printf("val:%d\n", n.val)
}

func main() {
	list := List{}
	list.Add(1)
	list.Add(2)
	sum := list.Sum()
	fmt.Println("list sum:", sum)

	list.travel(travelInt)
}

  