// go is a compiled language so it needs to be compiled first and then executed
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string = "Hello Tim"
	var number uint16 = 260 //explicit declaration
	var number2 = 260       //implicit declaration (khud se smajh lega type)
	number5 := true         //it will self declare the var and also assign it the value ("also implicit")
	fmt.Println("Hello betee!", name, number, number2, number5)
	fmt.Printf("%t", true)
	fmt.Printf("%e", 2.36347383738383838)
	fmt.Printf("%s", "hulloooo kiddooo")

	//the variables in go are static (i.e. unlike javascript they can store only one kind of data...not other)

	// by default int is always assigned as 0
	// by default bool is always assigned false
	// %T tells you the type of data
	// %v -value in default format
	//%t - tells you boolean
	//%% - literal %
	// %b - binary representation
	//%o - octal
	//%x - hexadecimal
	//%s - string
	//%e - scientificnotation
	//%g - for large exponents
	// %q - double quoted string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("type you birth year")
	scanner.Scan()                                       // to read input
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //base 10 number, size is 64 bits
	// strconv.ParseInt may return an error and '_' says pretty much ignore that error, can also write 'err' instead of '_'
	// if there's an error input string will be empty means 2021-0 =2021
	fmt.Printf("Ypu will be %d", 2021-input) //here initially input is a string hence its gonna give an error, thats why we used strconv.ParesInt

}

// go run command does the same thing for us, i.e. it compiles and then executes the code as well

// using go build you can make a compiled version of that code, which can be executed in any OS...even when you don't have go installed!
