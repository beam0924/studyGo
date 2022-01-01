package main

import (
	"fmt"
)

var (
	coins = 500
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, name := range users {
		count := 0
		for _, c := range name {
			switch {
			case c == 'e' || c == 'E':
				count += 1
				coins -= 1
			case c == 'i' || c == 'I':
				count += 2
				coins -= 2
			case c == 'o' || c == 'O':
				count += 3
				coins -= 3
			case c == 'u' || c == 'U':
				count += 4
				coins -= 4
			}
		}
		distribution[name] = count

	}
	return coins

}

func main() {

	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)

}
