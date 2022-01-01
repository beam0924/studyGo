package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("beam", 24)
	p2 := newPerson("wang", 25)
	fmt.Println(p1, p2)
}
