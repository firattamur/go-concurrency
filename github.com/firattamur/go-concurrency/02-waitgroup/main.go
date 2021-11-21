package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// counter for wait group
	var waitGroup sync.WaitGroup

	// increase counter
	// one go routine to wait for
	waitGroup.Add(1)

	// create anonymous function
	// call your function
	// decrease counter because routine is done
	go func() {
		repeat("Go")
		waitGroup.Done()
	}()

	// this wait until count in the wait group is zero!
	waitGroup.Wait()

}

func repeat(word string) {

	for i := 0; i <= 5; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * 500)
	}

}
