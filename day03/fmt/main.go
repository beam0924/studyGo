package main

import "fmt"

func main() {
	fmt.Print("长沙")
	fmt.Print("腾讯")
	fmt.Println()
	fmt.Println("----------")
	fmt.Println("长沙")
	fmt.Println("腾讯")
	fmt.Println("----------")
	//fmt.printf 格式化打印
	var m1 = make(map[string]int, 10)
	m1["长沙"] = 1
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)
	baifenbi := 90
	fmt.Printf("%d%%\n", baifenbi)
	fmt.Printf("%t\n", true)
	//整数->字符
	fmt.Printf("%q\n", 65)
	//浮点数和复数
	fmt.Printf("%b\n", 3.1415926535)
	//字符串
	fmt.Printf("%q\n", "海能卑下众水归") //字符串双引号表示
	fmt.Printf("%5.2s\n", "海能卑下众水归")
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是：", s)
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
}
