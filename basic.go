package main

import (
	"fmt"
	"time"
)

func attack(s string) {
	fmt.Println(s)
	time.Sleep(time.Second)
}

func main() {
	//define a slice

	items := []string{"Apple", "banana", "Strawberry", "olives"}
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()
	// this is a serial approach where each data is processed one by one..
	for _, item := range items {
		//method 1 - attack(item)
		go attack(item) // this is using go routine m-2
	}
	// the go attack are the
	time.Sleep(time.Second * 5)
}
