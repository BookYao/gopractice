/**
 * @Author: BookYao
 * @Description:
 update: 倒计时发射，但是回车可以打断
 * @File:  launch
 * @Version: 1.0.0
 * @Date: 2020/9/17 23:36
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launch...")
}

func main() {
	tick := time.Tick(1 * time.Second)
	/*for i := 10; i > 0; i-- {
		fmt.Println(i)
		<-tick
	}*/

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for countDown := 10 ; countDown > 0; countDown -- {
		fmt.Println(countDown)
		select {
		//case <-time.After(10 * time.Second):
		case <- tick:
		case <- abort:
			fmt.Println("launch abort...")
			return
		}
	}

	launch()
}
