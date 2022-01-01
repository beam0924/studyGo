package main

import "fmt"

//for循环
func main() {
	//基本格式
	for i := 10; i < 15; i++ {
		fmt.Println(i)
	}

	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	//变种2
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//无限循环
	//for {
	//fmt.Println("123")
	//}

	//for range循环
	s := "Hello腾讯"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
