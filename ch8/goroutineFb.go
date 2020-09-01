/**
 * @Author: BookYao
 * @Description: goruntine demo, calc fibonacci
 * @File:  goroutineFb
 * @Version: 1.0.0
 * @Date: 2020/9/1 23:02
 */

package main

import (
	"fmt"
	"time"
)

func waitTime(t time.Duration) {
	for {
		for _, ch := range `-\|/` {
			fmt.Printf("\r%c", ch)
			time.Sleep(t)
		}
	}
}

func fb(n int) int {
	if n < 2 {
		return n
	}

	return fb(n - 1) + fb(n - 2)
}

func main() {
	go waitTime(100 * time.Millisecond)
	const n = 45
	value := fb(n)
	fmt.Printf("Value:%d\n", value)
}

  