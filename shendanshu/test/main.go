package main

import (
	"fmt"
	"time"
)

func main() {
	rotMsgInterval := time.Minute * 60
	exp := fmt.Sprintf("0 */%d * * * *", int(rotMsgInterval.Minutes()))
	fmt.Println(exp)

}
