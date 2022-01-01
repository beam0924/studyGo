package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有初始化，没有在内存中开辟空间
	m1 = make(map[string]int, 10) //要估算该map的容量，避免程序运行期间动态扩容，不然影响执行效率
	m1["理想"] = 18
	m1["姬无命"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["娜扎"]) //如果不存在这个key，拿到对应类型的零值
	v, ok := m1["娜扎"]     //约定成俗，用ok接收返回bool值
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)
	}

	//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	delete(m1, "姬无命")
	fmt.Println(m1)
	delete(m1, "沙河") //删除不存在的key，就不做操作

}
