/**
 * @Author: BookYao
 * @Description: 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可
以同时与多个clock服务器通信，从多服务器中读取时间，并且在一个表格中一次显示所有服
务传回的结果，类似于你在某些办公室里看到的时钟墙。如果你有地理学上分布式的服务器
可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，
这些实例监听不同的端口，假装自己在不同的时区。像下面这样：
$ TZ=US/Eastern ./clock2 -port 8010 &
$ TZ=Asia/Tokyo ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

执行结果： 返回不同时区的时间
 * @File:  areaTimeEcho
 * @Version: 1.0.0
 * @Date: 2020/9/2 22:34
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func timeEcho(conn net.Conn, area string) {
	defer conn.Close()
	timeLayout := "2006-01-02 15:04:05\n"
	for {
		loc, _ := time.LoadLocation(area)
		zoneTime := time.Now().In(loc)
		timeString := zoneTime.Format(timeLayout)

		//io.WriteString(conn, time.Now().Format("2006/01/02 13:04:05\n"))
		io.WriteString(conn, timeString)
		time.Sleep(1 * time.Second)
	}
}

func timeEchoHandle(addrinfo string) {
	str := strings.Split(addrinfo, "=")
	area := str[0]
	addr := str[1]
	fmt.Printf("zone:%s-addrinfo:%s\n", area, addr)

	var zone string
	zone, ok := timezoneTable[area]
	if !ok  {
		zoneDefault := "Asia/Shanghai"
		zone = zoneDefault
	}
	fmt.Println("real zone:", zone)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Listen failed. addrinfo:%s-err:%v\n", addrinfo, err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept failed. addrinfo:%s-err:%v\n", addrinfo, err)
			break
		}

		go timeEcho(conn, zone)
	}
}

var timezoneTable map[string]string

func timezoneTableInit() {
	timezoneTable = make(map[string]string)
	timezoneTable["Shanghai"] = "Asia/Shanghai"
	timezoneTable["NewYork"] = "America/New_York"    //纽约
	timezoneTable["Dubai"] = "Asia/Dubai"            //迪拜
	timezoneTable["Moscow"] = "Europe/Moscow"         //莫斯科
}

/* Usage: ./areaTimeEcho beijing=localhost:8003 NewYork=localhost:8004
同时返回北京，纽约的时间 */
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s area:port\n", os.Args[0])
		return
	}

	timezoneTableInit()

	for _, val := range os.Args[1:] {
		fmt.Println("start listen: ", val)
		go timeEchoHandle(val)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

  