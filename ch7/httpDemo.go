/**
 * @Author: BookYao
 * @Description: 练习 7.11： 增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。例如，一
个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货
品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）
 * @File:  httpDemo
 * @Version: 1.0.0
 * @Date: 2020/8/28 10:30
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type dollers float32
type database map[string]dollers
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

/* Usage: http://192.168.50.132:8003/list
or http://192.168.50.132:8003/price?item=socks
or http://192.168.50.132:8003/help
*/
/*func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %.2f\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Not Such Item:%q", item)
			return
		}
		fmt.Fprintf(w, "price: %.2f", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Such Page:%s", r.URL)
		return
	}
}*/

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

func (db database) List(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %.2f\n", item, price)
	}
}

func (db database) Price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Such Item:%q", item)
		return
	}
	fmt.Fprintf(w, "Price: %.2f\n", price)
}

func (db database) Update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	fmt.Fprintf(w, "return item:%q", item)
	price := r.URL.Query().Get("price")
	fmt.Fprintf(w, "return price:%q", price)
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Such Item:%q", item)
		return
	}
	priceFloat, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Price format failed:%q", price)
		return
	}
	db[item] = dollers(priceFloat)
	fmt.Fprintf(w, "New Price: %.2f\n", priceFloat)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	//mux := http.NewServeMux()
	/*mux.HandleFunc("/list", http.HandlerFunc(db.List))
	mux.HandleFunc("/price", http.HandlerFunc(db.Price))
	log.Fatal(http.ListenAndServe("192.168.50.132:8003", mux))*/

	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/update", db.Update)
	log.Fatal(http.ListenAndServe("192.168.50.132:8003", nil))
}

  