package main

import (
	"fmt"
	"time"
)

func main() {

	// create channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Faster 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Slower 2s"
			time.Sleep(2 * time.Second)
		}
	}()

	// receive messages from c1 and c2 and print them
	// in this way messages from c2 will have to wait the
	// channel c1 even its faster and ready to receive.
	// because receiving is also and blocking operation
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// solution
	for {

		select {

		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}

	}

}
