package main

import (
	"fmt"
	"sync"
)

var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBuffChannel() {
	fmt.Println(b) // nil
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b = make(chan int) // 不带缓冲区的通道初始化
	b <- 10            // 阻塞
	fmt.Println("10发送到了通道b中了...")
	wg.Wait()
}

func buffChannel() {
	fmt.Println(b)
	b = make(chan int, 1) // 带缓冲区的通道初始化
	b <- 10
	fmt.Println("10发送到了通道b中了...")
	x := <-b
	fmt.Println("后台goroutine从通道b中取到了", x)
	close(b)
}

func main() {
	noBuffChannel()
	buffChannel()
}
