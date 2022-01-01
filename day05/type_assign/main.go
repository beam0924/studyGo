package main

import "fmt"

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("类型猜错了")
	} else {
		fmt.Println(str)
	}
}

func assign_1(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("a is a strinig, value is %v\n", v)
	case int:
		fmt.Printf("a is an int, value is %v\n", v)
	case bool:
		fmt.Printf("a is a bool, value is %v\n", v)
	default:
		fmt.Println("Unsupport type!")
	}
}
func main() {
	assign(100)
	assign("tencent")
	assign_1(100)
	assign_1("tencent")

}
