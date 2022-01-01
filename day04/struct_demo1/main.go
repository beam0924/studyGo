package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "beam"
	p.age = 24
	p.gender = "男"
	p.hobby = []string{"guitar", "badminton", "coding"}
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	fmt.Println(p.name)

	var p2 person
	p2.name = "wang"
	fmt.Println(p2)

	var s struct {
		x int
		y string
	}
	s.x = 100
	s.y = "heiheihei"
	fmt.Printf("type:%T value:%v\n", s, s)
}
