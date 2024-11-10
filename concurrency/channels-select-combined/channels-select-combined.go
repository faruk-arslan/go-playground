package main

import (
	"fmt"
	"time"
)

func main() {
	tick := make(chan int)
	boom := make(chan int)

	go processTick(tick)
	go processBoom(boom)

	// We could use the time.Tick but used explicit channels to demonstrate
	/*
		tick := time.Tick(100 * time.Millisecond)
		boom := time.After(500 * time.Millisecond)
	*/

	for {
		select {
		// We didn't use the value from the channel,
		// so we don't need to assign it to a variable
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM")
			// Since we have the return statement here, we don't need to close the channel
			// If we don't return, the default case would run infinitely
			return
		default:
			fmt.Println("...")
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func processTick(ch chan int) {
	for {
		ch <- 1
		time.Sleep(200 * time.Millisecond)
	}

}

func processBoom(ch chan int) {
	for {
		time.Sleep(1000 * time.Millisecond)
		ch <- 1
	}

}
