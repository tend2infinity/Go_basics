// go routines continued
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} //read/write mutex, which is basically a lock . If mutex is locked and someone tries to manipulate the value, it has to wait unless its unlocked.

//RWMutex allows for either at least one reader or exactly one writer.

//RLock is a shared read lock. When a lock is taken with it, other threads* can also take their own lock with RLock. This means multiple threads* can read at the same time. It's semi-exclusive.

//If the mutex is read locked, a call to Lock is blocked**. If one or more readers hold a lock, you cannot write.

//If the mutex is write locked (with Lock), RLock will block**.

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2) //spawning 20 go routines in total
		m.RLock() // lock needs to bne before sub routine
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
