package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 5, 566, 4, 34, 3, 45}
	//range
	for i, element := range a {
		fmt.Printf("%d: %d \n", i, element)
	}
	for _, element := range a {
		fmt.Printf("%d:\n", element)
	}
	// _ means an anonymous variable, used just to prevent the variable declared but not used error

	//maps
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	// or it can be initialized as mp :=  map[string]int{
	// 	"apple":  5,
	// 	"pear":   6,
	// 	"orange": 9,
	// }

	delete(mp, "pear")
	val, ok := mp["apple"] // if apple key exists then its val is stored in "val" and ok will be true, else default value will get stored and ok will be false

	fmt.Println(val, ok)
	fmt.Println(mp)

}
