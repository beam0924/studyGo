package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "beam",
		Age:  18,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	} else {
		fmt.Printf("%v\n", string(b))
	}

	// 反序列化
	str := `{"name":"beam","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了函数json.Unmarshal内部可以修改p2的值
	fmt.Printf("%#v\n", p2)
}
