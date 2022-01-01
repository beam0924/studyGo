package main

import "fmt"

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n%p\n%p\n", &(m.a), &(m.b), &(m.c))
}
