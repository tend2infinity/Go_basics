// go subroutine with wg (wait group), so that we don't need to use time.sleep
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // it synchronise multiple subroutines together

func main() { //here wg adds 0, cuz main is also a subroutine
	var msg = "hello"
	wg.Add(1) //initially wg starts with 0, so we add 1
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // its gonna decrement the number of waiting grps, since we add 1, so now its gonna be 0
	}(msg)
	msg = "Goodbye" //now this has no effect cuz its decoupled from the above function
	wg.Wait()       // to exit the wait grp
}
