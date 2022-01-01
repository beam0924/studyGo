package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := `PATH:/Users/beam/Desktop/秋招`
	//字符串相关的操作
	//字符串长度打印
	fmt.Println(len(s1))

	//字符串拼接
	name := "理想"
	world := "大帅比"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	//字符串分割
	ret := strings.Split(s1, "/")
	fmt.Println(ret)

	//判断包含
	fmt.Println(strings.Contains(ss, "理想"))
	fmt.Println(strings.Contains(ss, "理性"))

	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(ss, "理想"))
	fmt.Println(strings.HasSuffix(ss, "理想"))

	//判断子串出现的位置
	s2 := "abcdeb"
	fmt.Println(strings.Index(s2, "c"))
	fmt.Println(strings.Index(s2, "b"))
	fmt.Println(strings.LastIndex(s2, "b"))

	//拼接操作
	fmt.Println(strings.Join(ret, "+"))
}
