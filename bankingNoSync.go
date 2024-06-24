package main

import (
	"fmt"
	"sync"
	"time"
)

// we will define 3  functions one to withdraw money and one to deposit money and one to display money......

// this is dealing with critical section  problem so while there is an addition there shall be no triggering of adding to account ,

var balance = 0

func Deposit(wg *sync.WaitGroup, money int) {

	time.Sleep(time.Second)
	balance = balance + money
	wg.Done()
}
func Withdraw(wg *sync.WaitGroup, money int) {

	time.Sleep(time.Second)
	balance = balance - money
	wg.Done()
}
func Status() {

	fmt.Println(balance)
}

func main() {
	//here we will have operation to add deposit and delete
	var w sync.WaitGroup

	for i := 0; i < 150; i++ {

		if i%2 == 0 {
			w.Add(1)
			go Deposit(&w, 500)
		} else {
			w.Add(1)
			go Withdraw(&w, 500)
		}
	}
	// Block until the WaitGroup counter
	// goes back to 0; all the workers
	// notified they’re done.
	w.Wait()
	fmt.Println("Value of x", balance)
}
