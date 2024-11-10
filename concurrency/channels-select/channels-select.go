package main

import (
	"fmt"
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
		}

	}
}

func printFactorialMultipliers(num int, ch chan int) {
	// If we don't close the channel, it will cause a deadlock
	// Because there is no default case in the select statement
	defer close(ch)
	if num <= 1 {
		ch <- 1
		return
	}
	for i := num; i > 0; i-- {
		ch <- i
	}
}
