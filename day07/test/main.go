package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch chan int
	ch = make(chan int)
	go func() {
		defer wg.Done()
		ch <- 1
	}()
	x := <-ch
	fmt.Println(x)

	wg.Wait()
}
