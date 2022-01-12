//channel cont.
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) { //receiving data from the channel only
		for i := range ch { //better practice than prev example
			fmt.Println(i)
		} //here this leads to aa deadlock cuz, we are continuously reading data, but sending only twice

		//another syntax for line 15 is
		//for {
		// if i,ok:= <-ch;ok{
		//print i
		//}else{
		// break
		//}
		//}
		wg.Done()

	}(ch)
	go func(ch chan<- int) { //sending data into the channel only
		ch <- 42
		ch <- 27
		close(ch) //if we do this then error at line 17 gets cancelled
		//we can't close a channel twice
		// if we close a channel, you will never be able to pass data to it again
		//closing a subroutine is important, either use defer statement to close your channel, so that your sub routine can end peacefully, or use select statement (similar to switch statement) with the help of something like an empty struct to make a case where our goroutine gets closed smoothly
		wg.Done()
	}(ch)
	wg.Wait()
}
