package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	ret := time.Unix(1621222373, 0)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Day())

	fmt.Println(time.Nanosecond)

	fmt.Println(now.Add(24 * time.Hour))

}
