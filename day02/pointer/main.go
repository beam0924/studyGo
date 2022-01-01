package main

import "fmt"

func main() {
	//1.取地址
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	//2, *:根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	/*
		var a *int //nil point
		fmt.Println(a)
		*a = 100
		fmt.Println(*a)
	*/
	//new函数申请一个内存地址
	var a = new(int)
	fmt.Println(a)
	fmt.Println(*a) //初始化零值
	*a = 100
	fmt.Println(*a)

	//var b map[string]int
	//b["tencent"] = 100
	//fmt.Println(b)
}
