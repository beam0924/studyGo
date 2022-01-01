package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1 //赋值  浅拷贝
	var a3 = make([]int, 3, 5)
	copy(a3, a1) //copy 深拷贝
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//从切片中删除元素
	a := []int{30, 31, 33, 34, 35, 36, 37}
	fmt.Println(a)
	//要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 34 35 36 37]

	var b = make([]int, 5, 10)
	fmt.Println(b)
	for i := 0; i < 10; i++ {
		b = append(b, i)
	}
	fmt.Println(b)

	var c = []int{3, 1, 5, 8, 4, 2}
	sort.Ints(c)
	fmt.Println(c)

}
