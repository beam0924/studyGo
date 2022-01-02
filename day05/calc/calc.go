package calc

import "fmt"

func init() {
	fmt.Println("import本包时，自动执行.")
}

func Add(x, y int) int {
	return x + y
}
