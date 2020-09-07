/**
 * @Author: BookYao
 * @Description:
 * @File:  channelLinkPrint
 * @Version: 1.0.0
 * @Date: 2020/9/7 16:46
 */

package main

import "fmt"

func channelPipeDemo1() {
	natuals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			natuals <- i
		}
		close(natuals)
	}()

	go func() {
		for x := range natuals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Printf("Result:%d\n", x)
	}
}

func natuals(out chan <-  int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squares(out chan <- int, in <- chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func channlePrint(in  <- chan int) {
	for val := range in {
		fmt.Println("val: ", val)
	}
}

func channelPipeDemo2() {
	natral := make(chan int)
	square := make(chan int)
	go natuals(natral)
	go squares(square, natral)
	channlePrint(square)
}

func main() {
	channelPipeDemo1()
	channelPipeDemo2()
}

  