package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)
	// saara kahaani c++ jaisi hai
	//'*' before a datatype means a reference variable is required
	//'*' before a variable name means dereferencing the variable to get its value

	//'&' means address to a variable

	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

}
