package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go printFactorialMultipliers(5, ch)

	// select statement is used to wait on multiple communication operations
	// The select statement blocks until one of its cases can run, then it executes that case
	// If multiple cases can run, it chooses one at random
	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Println(i)
			} else {
				return
			}
		default:
			fmt.Println("default case")
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func printFactorialMultipliers(num int, ch chan int) {
	// We don't need t close the channel here, it won't create a deadlock
	// Because we have a default case in the select statement
	// But the default case will run continiously
	// If any other case has a return statement based ons a condition
	// It will stop liseing to the channel
	// We could have done that by closing the channel or listenin another channel and return
	// When we receive a message from that channel
	if num <= 1 {
		ch <- 1
		return
	}
	for i := num; i > 0; i-- {
		ch <- i
	}
}
