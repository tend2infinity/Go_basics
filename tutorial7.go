package main

import "fmt"

func main() {
	//muatble and immutable
	var x int = 5 //int is immutable
	y := x
	y = 7
	fmt.Println(x, y)

	var x1 []int = []int{3, 4, 5} //slice is mutable
	//simlarly map is mutable
	y1 := x1
	y1[0] = 100
	fmt.Println(x1, y1)
	//here x1 and y1 will be same

	var x2 [2]int = [2]int{3, 4} //arrays are also mutable
	y2 := x2
	y2[0] = 100
	fmt.Println(x2, y2)
	// but here x1 and y1 won't be the same

	//note : in case of slice, yaha bhi address pass hota hai
}
