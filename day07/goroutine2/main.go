package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记减一
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(300)))
	fmt.Println("Hello goroutine:", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //启动一个goroutine就加一
		go hello(i)
	}
	wg.Wait() //等待所有goroutine都结束
}
