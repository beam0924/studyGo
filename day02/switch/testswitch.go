package main

import "fmt"

func testSwitch() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)

	}
}

func testSwitch2() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age >= 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

// func main() {
// 	testSwitch()
// 	testSwitch2()

// }
