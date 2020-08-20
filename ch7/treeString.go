/**
 * @Author: BookYao
 * @Description:
练习 7.3： 为在gopl.io/ch4/treesort (§4.4)的*tree类型实现一个String方法去展示tree类型的值
序列
 * @File:  treeString
 * @Version: 1.0.0
 * @Date: 2020/8/20 16:17
 */

package main

import "fmt"

type tree struct  {
	value int
	left, right *tree
}

func treeAdd(t *tree, val int) *tree {
	if t == nil {
		t = new(tree)
		t.value = val
		return t
	}

	if val < t.value {
		t.left = treeAdd(t.left, val)
	} else {
		t.right = treeAdd(t.right, val)
	}

	return t
}

func treeTravel(t *tree) {
	if t == nil {
		return
	}

	treeTravel(t.left)
	fmt.Println("value:", t.value)
	treeTravel(t.right)
}

func (t *tree) String() string {
	s := ""
	if t == nil {
		return s
	}

	s += t.left.String()
	s += fmt.Sprintf("%d ", t.value)
	s += t.right.String()

	return s
}

func main() {
	t := treeAdd(nil, 5)
	t = treeAdd(t, 4)
	t = treeAdd(t, 6)
	t = treeAdd(t, 2)

	treeTravel(t)
	s := t.String()
	fmt.Println("Tree String:", s)
}
  