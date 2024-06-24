package main

import (
	"fmt"
	"sync"
)

// we will define 3  functions one to withdraw money and one to deposit money and one to display money......

// this is dealing with critical section  problem so while there is an addition there shall be no triggering of adding to account ,

var balance = 0

func deposit(wg *sync.WaitGroup, m *sync.Mutex, money int) {

	m.Lock()
	balance = balance + money
	wg.Done()
	m.Unlock()
}
func withdraw(wg *sync.WaitGroup, m *sync.Mutex, money int) {

	m.Lock()
	balance = balance - money
	wg.Done()
	m.Unlock()
}
func status() {

	fmt.Println(balance)
}

func main() {
	//here we will have operation to add deposit and delete
	var w sync.WaitGroup

	// This mutex will synchronize access to state.
	var m sync.Mutex

	for i := 0; i < 1000; i++ {

		if i%2 == 0 {
			w.Add(1)
			go deposit(&w, &m, 500)
		} else {
			w.Add(1)
			go withdraw(&w, &m, 500)
		}
	}
	// Block until the WaitGroup counter
	// goes back to 0; all the workers
	// notified they’re done.
	w.Wait()
	fmt.Println("Value of x", balance)
}
