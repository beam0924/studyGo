package main

import "fmt"

func f(n uint64) uint64 {
	//计算n的阶乘
	if n == 1 {
		return 1
	} else {
		return n * f(n-1)
	}
}

func main() {
	ret := f(10)
	fmt.Println(ret)
}
