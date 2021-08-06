package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	/*
		拷贝文件：
	*/
	srcFile := "/Users/didi/Desktop/Recent-Notes/Go入门笔记/第一阶段源码/14_IO操作/l_file/gopher.jpeg"
	destFile := "tomato.jpeg"
	//total,err := CopyFile1(srcFile,destFile)
	// total,err := CopyFile2(srcFile,destFile)
	// total, err := CopyFile3(srcFile, destFile)
	total, err := CopyFile4(srcFile, destFile)
	fmt.Println(total, err)
}

func CopyFile3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	// os.ModePerm 或者 0777
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

func CopyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	// 文件对象需要关闭 谨记
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}

//该函数：用于通过io操作实现文件的拷贝，返回值是拷贝的总数量(字节),错误
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕。。")
			break
		} else if err != nil {
			fmt.Println("报错了。。")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}


// 手写
func CopyFile4(srcFile, destFile string) (int, error) {
	// 打开文件
	// 读数据文件对象
	file1, err := os.Open(srcFile)
	if err != nil {
		// fmt.Println(err)
		return 0, err
	}
	// 写数据文件对象
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	// 关闭文件先行
	defer file1.Close()
	defer file2.Close()

	// 缓冲区
	bs := make([]byte, 1024, 1024)
	// 记录写入的字节数
	total := 0
	// 初始化 n1，读取的数据量
	n1 := -1
	for {
		// 写数据到缓存区
		n1, err = file1.Read(bs)
		if err == io.EOF || n1 == 0 {
			fmt.Println("拷贝完毕。。。")
			break
		} else if err != nil {
			fmt.Println("出错了")
			return total, err
		}
		total += n1
		// 从缓冲区写入数据
		n2, err := file2.Write(bs[:n1])
		if err != nil {
			fmt.Println(err)
			// 去除写入失败的数据字节数
			total = total - (n1 - n2)
			break
		}
	}
	return total, nil
}