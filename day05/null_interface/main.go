package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	m1 := make(map[string]interface{}, 20)
	m1["name"] = "beam"
	m1["age"] = 18
	m1["married"] = false
	m1["hobby"] = [...]string{"sing", "dance", "rap", "basketball"}
	fmt.Println(m1)

	s1 := []interface{}{"beam", 19, 19.01, false, [...]string{"1", "2", "3"}}
	fmt.Println(s1)
	s1 = append(s1, 0)
	fmt.Println(s1)
	show(false)
	show(nil)
	show(m1)
}
