package main

import (
	cc "beamwang/day05/calc"
	"fmt"
)

func init() {
	fmt.Println("自动执行.")
}

func main() {
	a := cc.Add(1, 3)
	fmt.Println(a)
}
