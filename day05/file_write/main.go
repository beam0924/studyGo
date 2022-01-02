package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	file_Obj, err := os.OpenFile("./xx.txt", os.O_TRUNC|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
	defer file_Obj.Close()
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	file_Obj.Write([]byte("welcome to tencent!\n"))
	file_Obj.WriteString("Good luck!\n")
}

func writeDemo2() {
	file_Obj, err := os.OpenFile("./xxx.txt", os.O_TRUNC|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0744)
	defer file_Obj.Close()
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	wr := bufio.NewWriter(file_Obj)
	wr.WriteString("Hello沙河\n")
	wr.Flush()
}

func writeDemo3() {
	str := "Hello, future"
	err := ioutil.WriteFile("./xxxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	writeDemo1()
	writeDemo2()
	writeDemo3()
}
