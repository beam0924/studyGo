package main

import "fmt"

type person struct {
	name   string
	gender string
}

func f(x person) {
	x.gender = "女" //修改的是副本的gender
}
func g(x *person) {
	x.gender = "女" //修改的就是传入的gender，可以直接用指针操作（Go语言语法糖）
	// (*x).gender = "女"
}
func main() {
	var p person
	p.name = "周林"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
	g(&p)
	fmt.Println(p.gender)

	p2 := person{
		name:   "beam",
		gender: "男",
	}
	fmt.Printf("%#v\n", p2)

	p3 := person{
		"wang",
		"男",
	}
	fmt.Printf("%#v\n", p3)

}
