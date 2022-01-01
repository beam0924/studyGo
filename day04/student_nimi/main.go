package main

import "fmt"

type person struct {
	string
	int
}

func main() {
	p1 := person{"beam", 18}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
