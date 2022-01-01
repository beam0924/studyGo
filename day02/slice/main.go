package main

import "fmt"

func main() {
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "深圳"}
	fmt.Println(s1, s2)

	//长度和容量
	fmt.Println(len(s1), len(s2))
	fmt.Println(cap(s1), cap(s2))

	//2，由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于一个数组切割，左闭右开
	fmt.Println(s3)
	s4 := a1[:4]
	s5 := a1[3:]
	s6 := a1[:]
	fmt.Println(s4, s5, s6)

	//切片的容量是指底层数组的容量
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
	//底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	//切片再切片
	s8 := s5[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))

	//切片是一个引用类型，他们都指向了底层的一个数组
	a1[6] = 1200 //修改了底层数组的值
	fmt.Println("s5", s5)
	fmt.Println("s8", s8)

}
