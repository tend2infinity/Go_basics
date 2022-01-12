// go rountines

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // it synchronise multiple subroutines together

func main() {
	go sayHello() //here this won't show anything cuz main() itself is a goroutine and when we spawn a goroutine inside of another one the inside one doesn't get time to get executed cuz the outer one dies instantly and to get around this we use (time.Sleep)

	// gyaaaan
	// instead of creating heavy overhead threads, we create an abstraction of a thread called go routine
	// inside of go runtime, we have a scheduler that maps these go routines to operating system threads for periods of time. it also assigns a certain amount of time to complete a particular process for a particular thread.
	// go rountines can start with very small stack spaces, cuz they can be reallocated very quickly and so they are very cheap to create and destroy
	// a simple go app can have thousands od go routines

	time.Sleep(100 * time.Millisecond) // before getting to the sleep part, even if our program encounters another subroutine, it never executes it, only after reaching here it executes the inner one, but it's not a good practice!

	//another example

	var msg = "namaste"
	go func() {
		fmt.Println(msg)
	}() // here this will print namaste cuz we passed it as a parameter, but if don't pass it , it will print "goodbye" since go routine get executed when we call time.sleep (similar to race condition)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println(("Hello"))
}

// to check race condition use go run -race tutorial12.go
