package main

import "fmt"

//函数的定义
func sum(i int, j int) int {
	return i + j
}

// 没有返回值的函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数也没有返回值的函数
func f2() {
	fmt.Println("f2")
}

//没有参数但是有返回值的函数
func f3() int {
	return 3
}

//命名返回值就当于提前声明一个变量，并且不需要显式返回
func f4(x int, y int) (res int) {
	res = x + y
	return
}

//多个返回值的函数
func f5() (int, int) {
	return 1, 2
}

//参数的类型简写(当参数中连续多个参数的类型一致时，我们可以将非最后一个的参数类型省略)
func f6(x, y int, m, n string, i, j bool) int {
	return x + y
}

//可变长参数的函数，可变长的参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y是int类型的切片（...后面指定的）
}

func main() {
	fmt.Println(sum(3, 4))
	i, j := f5()
	fmt.Println(i, j)
}
