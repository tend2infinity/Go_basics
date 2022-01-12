package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 100
	fmt.Println(arr)
	fmt.Println(arr[0])
	arr2 := [3]int{4, 5, 6}
	fmt.Println(len(arr2))
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Println((sum))
	//slice - its just a part of an array, imp part - it will have a length and a capacity where capacity is the number of elements that can be accomodated inside that slice after reslicing
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3] //start at 1 till 3 index, but do not include 3
	fmt.Println(len(s))
	fmt.Println(cap(s))     //capacity
	fmt.Println(s[:cap(s)]) //reslicing of slice
	fmt.Println(s)

	var a []int = []int{5, 6, 7, 8, 9} // declaration of a new slice (underlying principle is it declares an array and then take a slice of it)
	fmt.Println(cap(a))
	b := append(a, 10)
	fmt.Println(b)
	//or
	a = append(a, 10) // this chances the contents of slice a
	fmt.Println(a)
	a1 := make([]int, 5) // another way to make a slice
	fmt.Println(a1)
}
