/**
 * @Author: BookYao
 * @Description:
 * @File:  methodMap
 * @Version: 1.0.0
 * @Date: 2020/8/16 22:48
 */

package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}

	return ""
}

func (v Values) Add(key string, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := make(Values)

	m.Add("item1", "first")
	m.Add("item2", "second-2")
	m.Add("item2", "second")


	fmt.Println("item1:", m.Get("item1"))
	fmt.Println("item2:", m.Get("item2"))
}

  