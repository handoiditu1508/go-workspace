package sample

import "fmt"

func RunChannelSample() {
	fmt.Println("simple channel")
	// define a channel
	c := make(chan int)
	// var c chan int = make(chan int)
	// run a function in background
	go func() {
		fmt.Println("goroutine process")
		c <- 10 // write data to a channel
		c <- 20
		// close(c)
	}()

	go func() {
		fmt.Println("another goroutine process")
		val1 := <-c
		fmt.Printf("another goroutine read from channel: %d\n", val1)
	}()

	fmt.Println("no more routine")

	val, ok := <-c // read data from a channel
	fmt.Printf("value: %d\n", val)
	if !ok {
		fmt.Println("channel closed")
	} else {
		fmt.Println("channel still open")
	}
}
