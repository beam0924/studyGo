package main

import "fmt"

var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	// 函数中查找变量的顺序：
	// 1.先在函数内部查找
	// 2.找不到就往函数的外面查找，一直找到全局
	fmt.Println(x)
}
func main() {
	f1()
}
