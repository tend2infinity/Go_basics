package main

import (
	"fmt"
)

func main() {
	x := 3
	for x < 5 {
		fmt.Println(x)
		x++
	}
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}
	// break and continue works as it is
	ans := -1
	//switch statement
	switch ans {
	case 1, -1: // means when ans = 1
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not a case")
	}
}
