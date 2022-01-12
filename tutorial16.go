//channels
//concurrency and parallelism explained;
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // this channel send and receive only int data type
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() { //receiving go routine
			i := <-ch // arrow shows direction of flow of data
			fmt.Println(i)
			wg.Done()
		}()
		go func() { // sending go routine
			ch <- 42
			wg.Done()
		}()
	}

	wg.Wait()
}
