package main

import (
	"fmt"
)

func main() {
	var num1 float64 = 9
	var num2 int = 4
	answer := (num1) / float64((num2)) //operations happen on two same data types
	//follow BEDMASS Brackets->exponential->division.....
	// we can also compare strings, it compare letter by letter by its ASCII value
	fmt.Printf("%g", answer)
	// conditional <, > , <= , >= , == , !=
	x := 5
	//y:= 6
	val := x < 5
	fmt.Printf("%t \n", val)
	age := 12
	if age >= 18 {
		fmt.Println("You can ride akele")
	} else if age >= 14 {
		fmt.Println("You can ride with a parent")
	} else {
		fmt.Println("You cannot ride")
	}

}
