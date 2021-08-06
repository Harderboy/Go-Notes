package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/**
	遍历文件夹：
	*/
	dirname := "/Users/didi/Desktop/Recent-Notes/Go入门笔记/第一阶段源码/14_IO操作"
	//listFiles(dirname,0)
	traverseDir(dirname, 0)

}
func listFiles(dirname string, level int) {
	//level用来记录当前递归的层次，生成带有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//递归调用方法
			listFiles(filename, level+1)
		}
	}
}

func traverseDir(dirname string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := dirname + "/" + file.Name()
		fmt.Println(s + filename)
		if file.IsDir() {
			traverseDir(filename, level+1)
		}
	}
}
