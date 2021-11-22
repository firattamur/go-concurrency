package main

import (
	"fmt"
	"time"
)

func main() {

	// create a channel
	pipe := make(chan string)

	// call function and give word and pipe as argument
	go repeat("Go", pipe)

	// this code will receive only single message
	// to receive message from channel
	// msg := <-pipe
	// fmt.Println(msg)

	// to receive all messages from channel
	// but...
	// this will cause deadlock in the code
	// you will see fatal error: deadlock message in terminal
	// why?
	// because even the repeat function is done
	// we are waiting messages from the channel :/
	// solution for that is closing the channel
	// and not giving hope to channel listener any more
	// for {
	// 	// receive channel status with the message
	// 	// if channel is closed then why wait for?
	// 	msg, isOpen := <-pipe

	// 	if !isOpen {
	// 		break
	// 	}

	// 	fmt.Println(msg)
	// }

	// more clean way to solve channel close check
	// receive messages until channel is closed
	for msg := range pipe {
		fmt.Println(msg)
	}

}

// channel has type of 'chan'
// we need to specify data in the channel when we define a channel
func repeat(word string, pipe chan string) {

	for i := 0; i <= 5; i++ {
		// send word string though the channel
		// sending though the channel is blocking process!
		pipe <- word
		time.Sleep(time.Millisecond * 500)
	}

	// close the channel to prevent deadlock in the code
	// be carefull sender should close the channel not the receiver
	close(pipe)

}
