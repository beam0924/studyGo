package main

import "fmt"

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	// addr address
	address // 匿名嵌套结构体
}

type company struct {
	name string
	// addr address
	address
}

func main() {
	p1 := person{
		name: "beam",
		age:  18,
		address: address{
			province: "anhui",
			city:     "wuhu",
		},
	}
	fmt.Println(p1)
	// fmt.Println(p1.name, p1.address.city) // 非匿名嵌套结构体写法
	fmt.Println(p1.name, p1.city) // 匿名嵌套结构体写法 先在自己的结构体字段里面查找，找不到再去匿名嵌套的结构体中查找
}
