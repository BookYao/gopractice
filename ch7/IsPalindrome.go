/**
 * @Author: BookYao
 * @Description: 练习 7.10： sort.Interface类型也可以适用在其它地方。编写一个IsPalindrome(s
sort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。假
设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。
 * @File:  IsPalindrome
 * @Version: 1.0.0
 * @Date: 2020/8/27 17:26
 */

package main

import (
	"fmt"
	"sort"
)

type StringSlice string

func (s StringSlice) Len() int {
	return len(s)
}
func (s StringSlice) Swap(i, j int) {
	sTmp := []byte(s)
	sTmp[i], sTmp[j] = sTmp[j], sTmp[i]
}
func (s StringSlice) Less(i, j int) bool {
	sTmp := []byte(s)
	sTmp[i], sTmp[j] = sTmp[j], sTmp[i]
	return sTmp[i] == sTmp[j]
}

func isPalindrome(s sort.Interface) bool {
	fmt.Println("s.Len:", s.Len())
	i, j := 0, s.Len() - 1
	for j > i {
		if s.Less(i, j) == true {
			fmt.Println(i, j)
			i++
			j--
		} else {
			fmt.Println("222")
			return false
		}
	}

	return true
}

func stringReverse(s string) string {
	sLen := len(s)
	var sSlice []byte
	for i := 0; i < sLen; i++ {
		sSlice = append(sSlice, s[sLen - i - 1])
	}

	return string(sSlice)
}

func stringReverse2(s string) bool {
	sLen := len(s)
	for i, j := 0, sLen - 1; i < j; i, j = i + 1, j - 1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
func main() {
	var s StringSlice = "goog"
	fmt.Println("s:", s)
	fmt.Println("s[0]:", s[0])
	if isPalindrome(s) == true {
		fmt.Println("Is Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

	var str string = "googe"
	if str == stringReverse(str) {
		fmt.Println("String Reverse. Is Palindrome")
	} else  {
		fmt.Println("String Reverse. Not Palindrome")
	}

	if stringReverse2(str) == true {
		fmt.Println("Reverse2 Is Palindrome")
	} else  {
		fmt.Println("Reverse2, Not Palindrome")
	}
}

  