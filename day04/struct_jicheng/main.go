package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动!\n", a.name)

}

type dog struct {
	feet   uint8
	animal // animal拥有的方法，dog此时也有了
}

func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "beam",
		},
	}
	d1.wang()
	d1.move() // 这里模拟实现了继承
}
