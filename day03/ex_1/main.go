package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中汉字的数量
	//难点是判断一个字符时汉字
	s1 := "Hello沙河有沙안녕하세요"
	count := 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	//单词出现的次数统计
	s2 := "how do you do"
	lst := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range lst {
		// if _, ok := m1[w]; !ok { //这种写法ok变量作用域只在if语句中，节省了内存
		// 	m1[w] = 1
		// } else {
		// 	m1[w] += 1
		// }
		m1[w]++
	}
	fmt.Println(m1)
}
