package main

import "fmt"

func reset(board *[8][8]rune){
	// 数字自动会解引用，区别于切片
	board[0][0]='r'
}

func main(){
	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c\n",board[0][0])
	fmt.Printf("%v\n",board[0][0])
}