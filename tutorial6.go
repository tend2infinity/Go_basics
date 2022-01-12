package main

import "fmt"

// functions

func test(x int, y int) (int, int) {
	return x + y, x - y
}
func test2(x int, y int) (z1 int, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}

func test3(x int) {
	fmt.Println("hello!", x)
}

func test5(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() { //function closure
	return func() { fmt.Println(x) }
}

func main() {
	ans1, ans2 := test(5, 6)
	ans3, ans4 := test(5, 6)
	fmt.Println(ans1, ans2, ans3, ans4)
	x := test3 // similar to pointer to that function
	x(5)       // equal to test3()

	test4 := func() {
		fmt.Println("hellllloooooo!")
	}
	test4()

	add := func(x int) int {
		return x * -1
	}
	test5(add)
	returnFunc("namaste")() //isqa mtlb we need to return a function inside returnFunc()
}
