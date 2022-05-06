package main

import (
	"fmt"
	"sync"
	// "sync"
)

// var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	// defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	//<-ch1 单向通道报错
	close(ch1) //通道关闭后可读但不可写
}
func f2(ch1 chan int, ch2 chan int) {
	// defer wg.Done()
	t := 0
	for {
		i, ok := <-ch1 //当通道关闭时，ok = false
		t += 1
		if !ok {
			break
		}
		ch2 <- i * i
	}
	fmt.Println("!!!", t)
	once.Do(func() { close(ch2) })
}

func main() {
	var ch1 chan int
	var ch2 chan int
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	//defer close(ch2)
	// wg.Add(2)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	go f1(ch1)
	for i := range ch2 { // 在遍历时，如果channel没有关闭，则会出现deadlock的错误
		fmt.Println(i)
	}
	// wg.Wait()
}
