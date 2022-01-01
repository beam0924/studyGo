package main

import "fmt"

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}
type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动！")
}

func (c chicken) move() {
	fmt.Println("鸡动！")
}

func (c cat) eat(s string) {
	fmt.Printf("猫吃%s\n", s)
}

func (c chicken) eat(s string) {
	fmt.Printf("鸡吃%s\n", s)
}

func main() {
	var b animal
	fmt.Printf("%T\n", b)
	c1 := cat{"咪星人", 4}
	//c2 := chicken{2}
	b = c1
	fmt.Println(b)
	b.eat("fish")
	fmt.Printf("%T\n", b)

}
