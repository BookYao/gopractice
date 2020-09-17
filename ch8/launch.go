/**
 * @Author: BookYao
 * @Description:
 * @File:  launch
 * @Version: 1.0.0
 * @Date: 2020/9/17 23:36
 */

package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Launch...")
}

func main() {
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		<- tick
	}

	launch()
}

  