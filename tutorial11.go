//interfaces example 1
package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Writer interface {
	Write([]byte) (int, error) // function that takes a byte slice as argument and returns an integer and an error
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//interface example 2

type Incrementer interface {
	Increment() int
}
type IntCounter int //alias for int , IMP- we don't always need to have structs to implement interfaces, we can use any kind of custom type

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
