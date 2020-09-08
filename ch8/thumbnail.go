/**
 * @Author: BookYao
 * @Description:
  提取某个路径下的图片的缩略图
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

func main() {
	dirname := "./picDir"
	allPic, err := getAllPic(dirname)
	if err != nil {
		log.Printf("getAllPic failed. dirName:%s\n", dirname)
	}
	log.Println(allPic)
	makeThumbnail(allPic)
}

  