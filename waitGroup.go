package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func sim_cake(wg *sync.WaitGroup, cakeSignal chan int, i int) {
	defer wg.Done()

	waiting := rand.IntN(5)
	time.Sleep(time.Second * time.Duration(waiting))
	cakeCost := i
	cakeSignal <- cakeCost
}
func seq_cake(i int) {
	waiting := rand.IntN(5)
	time.Sleep(time.Second * time.Duration(waiting))
	fmt.Println(i)
}
func main() {

	cakeSignal := make(chan int, 3) // make a fifo channel of size 3
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sim_cake(&wg, cakeSignal, i)
	}

	// Close the channel once all goroutines are done this helps in running simultaneous go routine and wiating for all of them to complete.
	go func() {
		wg.Wait()
		close(cakeSignal)
	}()

	// Read from the channel until it is closed
	for message := range cakeSignal {
		fmt.Println(message)
	}

	defer func() {
		fmt.Println(time.Since(start))
	}()
}
