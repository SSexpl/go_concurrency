// Given the common fixed-size buffer, the task is to make sure that the producer can’t add data into the buffer
// when it is full and the consumer can’t remove data from an empty buffer. Accessing memory buffers should not be
// allowed to producers and consumers at the same time.

// 1. Define a fixed sized buffer size
package main

import (
	"fmt"
	"time"
)

const bufferSize = 5

func producer(buffer chan<- int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		buffer <- i // This will block if the buffer is full automatic block...
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}
	done <- true
}

func consumer(buffer <-chan int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		item := <-buffer // This will block if the buffer is empty but as the values are put the values will be consumed here.
		fmt.Printf("Consumed: %d\n", item)
		time.Sleep(time.Second) // Simulate some work
	}
	done <- true
}

func main() {
	buffer := make(chan int, bufferSize)
	doneProducer := make(chan bool)
	doneConsumer := make(chan bool)

	go producer(buffer, doneProducer)
	go consumer(buffer, doneConsumer)

	<-doneProducer
	<-doneConsumer
	fmt.Println("Processing complete")
}
