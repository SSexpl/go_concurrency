package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

var buffer = make([]int, 0, bufferSize)
var mutex = &sync.Mutex{}
var cond = sync.NewCond(mutex)

func producer(done chan<- bool) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		for len(buffer) == bufferSize {
			// Wait until there is space in the buffer
			cond.Wait()
		}
		buffer = append(buffer, i)
		fmt.Printf("Produced: %d\n", i)
		cond.Signal() // Notify a waiting consumer
		mutex.Unlock()
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}
	done <- true
}

func consumer(done chan<- bool) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		for len(buffer) == 0 {
			// Wait until there is something in the buffer
			cond.Wait()
		}
		item := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("Consumed: %d\n", item)
		cond.Signal() // Notify a waiting producer
		mutex.Unlock()
		time.Sleep(time.Second) // Simulate some work
	}
	done <- true
}

func main() {
	doneProducer := make(chan bool)
	doneConsumer := make(chan bool)

	go producer(doneProducer)
	go consumer(doneConsumer)

	<-doneProducer
	<-doneConsumer
	fmt.Println("Processing complete")
}
