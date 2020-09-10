/**
 * @Author: BookYao
 * @Description:
  提取某个路径下的图片的缩略图
  update1: add thumbnailVer2
  update2: add thumbnailVer3
  update3: add thumbnailVer4
 * @File:  thumbnail
 * @Version: 1.0.0
 * @Date: 2020/9/8 16:31
 */

package main

import (
	"fmt"
	"log"
	"os"
	"gopl.io/ch8/thumbnail"
	"sync"
)

func getAllPic(dirname string) (pic []string, err error){
	log.Printf("dirName:%s\n", dirname)
	if len(dirname) == 0 {
		log.Printf("Input dirname.\n")
		return nil, fmt.Errorf("Input DirName.")
	}

	f, err := os.OpenFile(dirname, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	//参数：n,表读取目录的成员个数。通常传-1,表读取目录所有文件对象
	fileinfo, _ := f.Readdir(-1)

	//获取系统的路径分隔符
	separator := string(os.PathSeparator)

	for  _, info := range fileinfo {
		if info.IsDir() {
			childDir := dirname + separator + info.Name()
			pic, err = getAllPic(childDir)
			if err != nil {
				log.Printf("Get ChildName Pic Failed. chileName:%s", childDir)
				continue
			}
		} else {
			//fmt.Println("infoName:", info.Name())
			fullName := dirname + separator + info.Name()
			pic = append(pic, fullName)
		}
	}

	//fmt.Printf("====== %s DirFile Get Finish...====\n", dirname)
	return pic, nil
}

func makeThumbnail(pic []string) {
	for _, f := range pic {
		_, err := thumbnail.ImageFile(f)
		if err != nil {
			log.Printf("make thumbnail failed. picName:%s-err:%v\n", f, err)
		}
	}
}

func makeThumbnailVer2(pic []string) {
	ch := make(chan struct{})
	for _, f := range pic {
		go func(f string) {
			if _, err := thumbnail.ImageFile(f); err != nil {
				log.Printf("make thumbnail ver2 failed. picName:%s-err:%v\n", f, err)
			}
			ch <- struct{}{}
		}(f)
	}

	for range pic {
		<- ch
	}

}

func makeTHumbnailVer3(pic []string) (thumbFiles []string, err error) {
	type item struct {
		thumbfile string
		err error
	}

	ch := make(chan item, len(pic))
	for _, f := range pic {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range pic {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}

		thumbFiles = append(thumbFiles, it.thumbfile)
	}

	return thumbFiles, nil
}

func makeThumbnailVer4(filename <- chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filename {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				fmt.Println("333")
				log.Printf("thumbnail failed. err:%v\n", err)
				return
			}
			fmt.Println("thumb:", thumb)
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	dirname := "./picDir"
	allPic, err := getAllPic(dirname)
	if err != nil {
		log.Printf("getAllPic failed. dirName:%s\n", dirname)
	}
	log.Println(allPic)
	//makeThumbnail(allPic)
	//makeThumbnailVer2(allPic)

	/*thumbFiles, err := makeTHumbnailVer3(allPic)
	if err != nil {
		fmt.Printf("make Thumbnail Ver3 failed. err:%v\n", err)
	}
	fmt.Println("ThumbFiles: ", thumbFiles)*/

	picChan := make(chan string, len(allPic))
	go func() {
		for _, pic := range allPic {
			picChan <- pic
		}
		close(picChan)
	}()

	total := makeThumbnailVer4(picChan)
	fmt.Printf("total:%d\n", total)
}

  