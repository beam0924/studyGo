package main

import "fmt"

func deferdome() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") //defer将该语句延迟到函数即将返回的时候再执行
	defer fmt.Println("哈哈哈") //一个函数中可以有多个defer语句
	defer fmt.Println("嘻嘻嘻") //多个defer语句按照先进先出顺序延时执行
	fmt.Println("end")
}

func main() {
	deferdome()
}
