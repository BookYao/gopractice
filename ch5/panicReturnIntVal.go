/**
 * @Author: BookYao
 * @Description: 练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
 * @File:  panicReturnIntVal
 * @Version: 1.0.0
 * @Date: 2020/8/12 23:35
 */

package main

import "fmt"

func paincReturnIntVal() (val int) {
	defer func() {
		if p := recover(); p != nil {
			val = p.(int)
		}
	}()
	panic(1)
}

func main() {
	val := paincReturnIntVal()
	fmt.Println("Int val:", val)
}

  