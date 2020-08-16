/**
 * @Author: BookYao
 * @Description:
 * @File:  cacheMethod
 * @Version: 1.0.0
 * @Date: 2020/8/16 23:57
 */

package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
} {
	mapping :make(map[string]string),
}

func getCache(key string) string {
	cache.Lock()
	defer cache.Unlock()
	v := cache.mapping[key]
	return v
}

func setCache(key, val string) {
	cache.Lock()
	defer cache.Unlock()
	cache.mapping[key] = val
}

func main() {
	setCache("item1", "first")
	setCache("item2", "second")
	fmt.Println("getItem1:", getCache("item1"))
	fmt.Println("getItem2:", getCache("item2"))
}

  