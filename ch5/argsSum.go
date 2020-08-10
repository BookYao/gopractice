/**
 * @Author: BookYao
 * @Description: 添加练习题 5.15 编写类似sum的可变参数函数max和min
 * @File:  argsSum
 * @Version: 1.0.0
 * @Date: 2020/8/10 17:28
 */

package main

import "fmt"

func sum(vals ...int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}
	return sum
}

func maxFunc(vals ...int) (int, error) {
	max := 0

	if len(vals) == 0 {
		return 0, fmt.Errorf("Max: %s", "至少传入一个参数")
	}

	for _, val := range vals {
		if (max <= val) {
			max = val
		}
	}
	return max, nil
}

func minFunc(vals ...int) (int, error) {
	min := 0

	if len(vals) == 0 {
		return 0, fmt.Errorf("Min: %s", "至少传入一个参数")
	}

	for _, val := range vals {
		if min >= val {
			min = val
		}
	}

	return min, nil
}

func main() {
	fmt.Printf("Sum: %d\n", sum(1, 2, 3, 4))

	max, err := maxFunc(1, 3, 2, 4, 8, 0, -1, 9)
	if err != nil {
		fmt.Printf("maxFunc Failed. %v\n", err)
	} else {
		fmt.Printf("Max:%d\n", max)
	}

	min, err := minFunc(1, 3, 2, 4, 8, 0, -1, 9)
	if err != nil {
		fmt.Printf("minFunc Failed. %v\n", err)
	} else  {
		fmt.Printf("Min:%d\n", min)
	}

	max, err = maxFunc()
	if err != nil {
		fmt.Printf("maxFunc Failed. %v\n", err)
	} else {
		fmt.Printf("Max:%d\n", max)
	}

	min, err = minFunc()
	if err != nil {
		fmt.Printf("minFunc Failed. %v\n", err)
	} else {
		fmt.Printf("Min:%d\n", min)
	}

}

  