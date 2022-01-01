package main

import "fmt"

func main() {
	s1 := "pig"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'b'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
	t := 'H'
	//i := 1
	var aa = 'w'
	fmt.Println(t, aa)
}
