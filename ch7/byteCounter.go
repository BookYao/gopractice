/**
 * @Author: BookYao
 * @Description: 增加 练习 7.1： 使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器
 * @File:  byteCounter
 * @Version: 1.0.0
 * @Date: 2020/8/19 13:49
 */

package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c = ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int
func (c *WordCounter) Write(p []byte) (n int, err error) {
	string := strings.NewReader(string(p))
	buf := bufio.NewScanner(string)
	buf.Split(bufio.ScanWords)

	sum := 0
	for buf.Scan() {
		sum++
	}

	*c = WordCounter(sum)

	return sum, nil
}

type LineCounter int
func (c *LineCounter) Write(s []byte) (n int, err error) {
	string := strings.NewReader(string(s))
	buf := bufio.NewScanner(string)
	buf.Split(bufio.ScanLines)

	line := 0
	for buf.Scan() {
		line++
	}

	*c = LineCounter(line)
	return line, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	fmt.Fprintf(&c, "%d", 123)
	fmt.Println(c)

	var word WordCounter
	word.Write([]byte("hello world! hello everyone!"))
	fmt.Println("word:", word)

	var line LineCounter
	line.Write([]byte("hello world!\n hello everyone!\n"))
	fmt.Println("line:", line)
}

  