package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func makeCakeofcost(cakeSignal chan int) {

	for i := 0; i < 3; i++ {
		//time.Sleep(time.Second)
		//	waiting := rand.IntN(5)
		//	time.Sleep(time.Second * time.Duration(waiting))
		cakeCost := i
		cakeSignal <- cakeCost //this will place the numbers in a
	}
	defer close(cakeSignal)
}
func sim_cake(cakeSignal chan int, i int) {

	waiting := rand.IntN(5)
	time.Sleep(time.Second * time.Duration(waiting))
	cakeCost := i
	cakeSignal <- cakeCost
}
func main() {

	cakeSignal := make(chan int, 3) // make a fifo channel of size 3

	//makeCakeofcost(cakeSignal)
	start := time.Now()
	go makeCakeofcost(cakeSignal)

	for {
		message, open := <-cakeSignal
		if !open {
			break
		}
		fmt.Println(message)
	}
	//	close(cakeSignal)
	defer func() {
		fmt.Println(time.Since(start))
	}()

}
