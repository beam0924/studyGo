package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动！")
}

func (c cat) eat(s string) {
	fmt.Printf("猫吃%s\n", s)
}

func main() {
	var a animal
	a = cat{"beam", 4}
	fmt.Println(a)
}
