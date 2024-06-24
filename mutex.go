// Program with race condition
package main

import (
	"fmt"
	"sync" // to import sync later on
)

var count = 0

// This is the function we’ll run in every
// goroutine. Note that a WaitGroup must
// be passed to functions by pointer.
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	count = count + 1
	// On return, notify the
	// WaitGroup that we’re done.
	m.Unlock()
	wg.Done()
}
func main() {

	// This WaitGroup is used to wait for
	// all the goroutines launched here to finish.
	var w sync.WaitGroup
	// This mutex will synchronize access to state.
	var m sync.Mutex
	// Launch several goroutines and increment
	// the WaitGroup counter for each
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w, &m)
	}
	// Block until the WaitGroup counter
	// goes back to 0; all the workers
	// notified they’re done.
	w.Wait()
	fmt.Println("Value of x", count)
}
