package main

import "fmt"

func main() {
	a := [...]int{1, 3, 5, 7, 8}
	count := 0
	for _, v := range a {
		count += v
	}
	fmt.Printf("The sum of a:%d\n", count)
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
