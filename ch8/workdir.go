/**
 * @Author: BookYao
 * @Description:
 * @File:  workdir
 * @Version: 1.0.0
 * @Date: 2020/9/20 0:43
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func directs(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dir Read Err:%v\n", err)
		return nil
	}
	return entries
}

func workdir(dir string, filesize chan <- int64) {
	for _, entry := range directs(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			workdir(subdir, filesize)
		} else  {
			filesize <-  entry.Size()
		}
	}
}

func printDisk(nfiles, nsize int64) {
	fmt.Printf("%d file. size:%.1f Mb\n", nfiles, float64(nsize)/1e6)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	filesize := make(chan int64)
	go func() {
		for _, root := range roots {
			workdir(root, filesize)
		}
		close(filesize)
	}()

	var nfiles, nsize int64
	for size := range filesize {
		nfiles++
		nsize += size
	}

	printDisk(nfiles, nsize)
}

  