package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	reflectSetValue(&a)
	fmt.Println(a)
}
