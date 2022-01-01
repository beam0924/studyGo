package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [4]bool
	//数组的长度是数组类型的一部分，所以这里的a1和a2既不可以比较也不可以赋值
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是零值(bool值就是false，整形和浮点型都是0，字符串就是“”)
	fmt.Println(a1, a2)
	//初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//初始化方式2
	//根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a10)
	//初始化方式3
	//根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//遍历方法1：根据索引来遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//遍历方法2：for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多维数组
	//[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b1[0] = 100
	fmt.Println(b1, b2)
}
