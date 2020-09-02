/**
 * @Author: BookYao
 * @Description:
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
	//timeLayout := "2006-01-02 15:04:05"
	for {
		io.WriteString(conn, time.Now().Format("2006/01/02 13:04:05\n"))
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Unix())
		/*utctime := time.Now().UTC()
		zone, _ := time.LoadLocation("UTC")
		tmp, _ := time.ParseInLocation(timeLayout, utctime, zone)
		timestamp := tmp.Unix()
		fmt.Println("timeStamp:", timestamp)*/
	}
}

func timeEchoHandle(addrinfo string) {
	str := strings.Split(addrinfo, "=")
	area := str[0]
	addr := str[1]
	fmt.Printf("zone:%s-addrinfo:%s\n", area, addr)
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

		go timeEcho(conn, area)
	}
}

var timezoneTable map[string]int

func timezoneTableInit() {
	timezoneTable = make(map[string]int)
	timezoneTable["Beijing"] = 8
	timezoneTable["Lodon"] = 0
	timezoneTable["New Delhi"] = 5
	timezoneTable["New York"] = -5
}

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

  