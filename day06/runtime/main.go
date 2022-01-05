package main

import (
	"fmt"
	"path"
	"runtime"
)

func f() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Call() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	f()
}
