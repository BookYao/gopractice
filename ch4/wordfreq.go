package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
count word count by os.stdin
*/
func wordfreq() {
	wordcount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordcount[input.Text()]++
	}

	for k, v := range wordcount {
		fmt.Printf("word: %s-count:%d\n", k, v)
	}
}

func main() {
	fmt.Println("Please input word:")
	wordfreq()
}
