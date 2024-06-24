package main

import (
	"fmt"
	"time"
)

func cake(is chan bool) {
	time.Sleep(time.Second)
	fmt.Println("cake is made")
	is <- true

}
func main() {

	start := time.Now()
	cakeSignal := make(chan bool)

	//order the go-rutine to make the cake
	go cake(cakeSignal) //cake() for-seq wait and then proceed

	//till that time start delivery operations
	fmt.Println("Start organising for delivery partner")
	time.Sleep(time.Millisecond * 500)
	//cakeSignal <- false
	//once the cake is made
	iscakeMade := <-cakeSignal //wait till you get a response on this.

	//do the delivery.
	if iscakeMade {
		fmt.Println("cake out for delivery")
	}
	defer func() {
		fmt.Println(time.Since(start))
	}()

}
