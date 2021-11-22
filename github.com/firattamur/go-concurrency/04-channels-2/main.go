package main

import "fmt"

func main() {

	// deadlock()
	nodeadlock()

}

func nodeadlock() {

	// give channel a capacity
	// in this way channel will not block until its full
	c := make(chan string, 2)
	c <- "Go - 1"
	c <- "Go - 2"

	// but ...
	// channel capacity is 2
	// if we add a third one then it will block
	c <- "Go - 3"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	// but... channel is full will be blocked
	msg = <-c
	fmt.Println(msg)

}

func deadlock() {

	c := make(chan string)
	c <- "Go"

	// we will not get the message from the channel
	// because sending though the channel is blocking operation
	// that's why we will not reach lines after the line 8
	msg := <-c
	fmt.Println(msg)

}
