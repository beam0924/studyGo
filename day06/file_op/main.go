package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./xxx.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	//因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed, err:%v\n", err)
		return
	}
	//读原文件，写入临时文件
	var ret [4]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}
	_, err = tmpFile.Write(ret[:n])
	if err != nil {
		fmt.Printf("write the tmpFile failed, err:%v\n", err)
	}

	//写入要插入的内容
	tmpFile.Write([]byte("强行加入一行\n"))

	//紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err = fileObj.Read(x[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read from rest file failed, err:%v\n", err)
			return
		}
		_, err = tmpFile.Write(x[:n])
		if err != nil {
			fmt.Printf("write the rest tmpFile failed, err:%v\n", err)
		}

	}

	//删除原文件，将临时文件改名为原文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./tmp.txt", "./xxx.txt")
}

func main() {
	f1()
}
