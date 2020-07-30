/**
 * @Author: BookYao
 * @Description:
 * @File:  fielEof
 * @Version: 1.0.0
 * @Date: 2020/7/30 16:26
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func testStdinEof() error {
	in := bufio.NewReader(os.Stdin)
	for  {
		r, _, err := in.ReadLine()
		if err == io.EOF {
			fmt.Println("Already read eof")
			break
		}

		fmt.Printf("Read str: %s\n", r)

		if err != nil {
			return fmt.Errorf("ReadRune failed, %v", err)
		}
	}

	return nil
}

func main() {
	fmt.Println("Read Line String. Please input string:")
	err := testStdinEof()
	if err != nil {
		log.Println("test stdin eof failed.")
	}
}

  