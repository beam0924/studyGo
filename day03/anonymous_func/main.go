package main

import "fmt"

func main() {
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 29)

	//如果只是调用一次的函数，我们还可以简写成立即执行函数
	func() {
		fmt.Println("Hello,沙河")
	}()
	func(a, b int) {
		fmt.Println(a + b)
		fmt.Println("Hello,沙河")
	}(100, 200)
}
