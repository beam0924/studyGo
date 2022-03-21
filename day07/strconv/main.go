package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 从字符串中解析出数字
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	//将数字转换成字符串类型
	i := int32(97)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)

	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v\n", ret3)
}
