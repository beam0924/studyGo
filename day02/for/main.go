package main

import "fmt"

//流程控制之跳出for循环
func main() {
	//当i=5时，跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("over")

	//当i=5时，跳过本次for循环（不执行for循环内部的打印语句，继续下一次循环）
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("over")
}
