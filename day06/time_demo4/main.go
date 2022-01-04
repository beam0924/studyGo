package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	// 按照对应的格式解析字符串
	t, err := time.Parse("2006/01/02 15:04:05", "2099/04/13 09:21:39")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(t)
	fmt.Println(t.Unix())
	time.Sleep(5 * time.Second)
	now1 := time.Now()
	d := now1.Sub(now)
	fmt.Println(d)
}
