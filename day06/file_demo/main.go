package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
	}
	fmt.Printf("%T\n", fileObj)
	fileInfo, err := fileObj.Stat()
	fmt.Printf("文件大小是：%d byte\n", fileInfo.Size())

	os.MkdirAll("./log/text.log", 0744)
	os.Rename("./log/text.log", "./log/text.log.bak")
}
