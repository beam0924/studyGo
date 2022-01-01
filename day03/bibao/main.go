package main

import "fmt"

func f1(f func()) {
	fmt.Println("This is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行包装
func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}

func main() {
	f1(f3(f2, 100, 200))
}
