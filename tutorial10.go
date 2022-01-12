//methods inside a struct

package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setAge(age int) {
	//method that acts on a student object whose name is s
	s.age = age

}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80}, 19}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1.age)
	s2 := Student{}
	s2.setAge(21)
	fmt.Println(s2)
}
