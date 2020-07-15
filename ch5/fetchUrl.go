package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s url", os.Args[0])
		os.Exit(0)
	}

	for _, url := range os.Args[1:] {
		fmt.Println("url:", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("http Get failed...")
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			log.Fatal("http read failed")
			os.Exit(0)
		}

		fmt.Printf("HTTP Url body: %s", b)
	}
}
