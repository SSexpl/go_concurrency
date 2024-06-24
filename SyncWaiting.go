package main

//this solves the problem of communication bw processes
//here there is a need to do the waiting in bw a thread hence we need an explicit communication between the processes.
// so if a wait is called in the
import (
	"fmt"
	"sync"
)

var balance = 0
var mu sync.Mutex
var cond = sync.NewCond(&mu)

func Deposit(wg *sync.WaitGroup, money int) {
	defer wg.Done()

	mu.Lock()
	balance = balance + money
	fmt.Printf("Deposited %d, new balance: %d\n", money, balance)
	cond.Signal()
	mu.Unlock()
}

func Withdraw(wg *sync.WaitGroup, money int) {
	defer wg.Done()

	mu.Lock()
	for balance < money {
		cond.Wait() // if the balance is not enough we will wait this wait will temporary release the mutex so that other competing process can start.
		//once the signal is called the control shifts back to this line.
	}
	balance = balance - money
	fmt.Printf("Withdrew %d, new balance: %d\n", money, balance)
	mu.Unlock()
}

func Status() {
	mu.Lock()
	fmt.Println("Current balance:", balance)
	mu.Unlock()
}

func main() {
	var w sync.WaitGroup

	for i := 0; i < 150; i++ {
		w.Add(1)
		if i%2 == 0 {
			go Deposit(&w, 500)
		} else {
			go Withdraw(&w, 500)
		}
	}

	w.Wait()
	Status()
	fmt.Println("All operations completed.")
}
