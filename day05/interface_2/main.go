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
type dog struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动！")
}

func (c cat) eat(s string) {
	fmt.Printf("猫吃%s\n", s)
}

func (d *dog) eat(s string) {
	fmt.Printf("狗吃%s\n", s)
}

func (d *dog) move() {
	fmt.Println("狗动！")
}

func main() {
	var a1 animal
	var a2 animal
	c1 := cat{"tom", 4}
	c2 := &cat{"tim", 4}
	a1 = c1
	a2 = c2
	fmt.Println(a1, a2)

	var a3 animal
	var a4 animal
	// d1 := dog{"tom", 4}
	d2 := &dog{"tim", 4}
	// a3 = d1  //这里会报错类型不匹配，接口不能接受值类型d1
	a4 = d2
	fmt.Println(a3, a4)
}
