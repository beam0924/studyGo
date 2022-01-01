package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n)
	//查看变量的值
	fmt.Printf("%v\n", n)
	//查看二进制
	fmt.Printf("%b\n", n)
	//查看十进制
	fmt.Printf("%d\n", n)
	//查看八进制
	fmt.Printf("%o\n", n)
	//查看十六进制
	fmt.Printf("%x\n", n)
	var s = "Hello, Tencent!"
	//查看字符串的值
	fmt.Printf("%s\n", s)
	//查看字符串的值（通用版）
	fmt.Printf("%v\n", s)
	//查看字符串的值（添加标识符）
	fmt.Printf("%#v\n", s)

}
