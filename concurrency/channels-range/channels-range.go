package main

import "fmt"

func main() {
	ch := make(chan int)

	go printFactorialMultipliers(5, ch)
	// v, ok := <-ch
	// The above syntax is used to check if the channel is closed or not
	// If the channel is closed, ok will be false
	for i := range ch {
		fmt.Println(i)
	}
}

func printFactorialMultipliers(num int, ch chan int) {
	// Only the sender side should close the channel
	// Sending on a closed channel will cause a panic
	// Channels are not like files, they usually don't need to be closed
	// But you have to close the channel if you are ranging over it
	defer close(ch)
	if num <= 1 {
		ch <- 1
		return
	}
	for i := num; i > 0; i-- {
		ch <- i
	}

}
