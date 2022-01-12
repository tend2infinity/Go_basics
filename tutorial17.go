//channel cont.
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) //by putting another parameter 50, we make it a buffered channel i.e. it will contain an internal data store, which is in this case 50 integers, so it will store any external values pushed to the channel , which we aren't taking out, so that our app doesn't reach a a deadlock, but a problem is, that particulatr extra message is lost (ref line 23)
	wg.Add(2)
	go func(ch <-chan int) { //receiving data from the channel only

		// specify in the arguments ki is subroutine me channel sending hai, ya receiving, if you don't do it then any subroute will send or receive data from the channel and program won't go right
		i := <-ch
		fmt.Println(i)
		wg.Done()

	}(ch)
	go func(ch chan<- int) { //sending data into the channel only
		ch <- 42
		ch <- 27 //we are sending two and receiving one....which will lead to a deadlock if we dont specify "50" above while making the channel
		wg.Done()
	}(ch)
	wg.Wait()
}
