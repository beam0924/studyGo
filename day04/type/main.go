package main

import "fmt"

type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {
	var n myInt = 100
	var m yourInt = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune // rune实际上就是int32的类型别名
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
