package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("Hello", i)
}

func main() {
	for i := 0; i < 20; i++ {
		go hello(i)
	}
	for j := 50; j < 70; j++ {
		go func(i int) {
			fmt.Println(i)
		}(j)
	}

	fmt.Println("main")
	time.Sleep(time.Second)
	// main函数结束了，由main函数启动的goroutine也都结束了
}
