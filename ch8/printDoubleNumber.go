/**
 * @Author: BookYao
 * @Description:
 * @File:  printDoubleNumber
 * @Version: 1.0.0
 * @Date: 2020/9/17 23:59
 */

package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	for i := 0; i < 10 ; i++ {
		select {
			case ch <- i:
			case x := <- ch:
				fmt.Println(x)
		}
	}
}

  