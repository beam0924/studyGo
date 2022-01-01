package main

import "fmt"

func main() {
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break //只跳出内层for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出外层for循环
		}
	}

	//goto + label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto LABEL //调到我们指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
LABEL: //label标签
	fmt.Println("over")
}
