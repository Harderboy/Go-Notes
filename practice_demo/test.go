package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sl := []byte{97, 98, 99, 100}
	s := string(sl)
	fmt.Println(s) //abcd

	sl1 := []byte(s)
	fmt.Println(sl1)

	var i rune = 'A'
	var j byte = 65
	// rune -> string
	s_i := string(i)
	fmt.Printf("%T-%#v-%v\n", s_i, s_i, s_i)
	// byte -> string
	s_j := string(j)
	fmt.Printf("%T-%#v-%v\n", s_j, s_j, s_j)

	r := strings.NewReader("some io.Reader stream to be read\n")
	// 注意写法 _
	// _ 不用使用，其他变量都必须被使用，比如这里的 err
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir("..")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
