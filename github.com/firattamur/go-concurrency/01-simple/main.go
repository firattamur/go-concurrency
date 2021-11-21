package main

import (
	"fmt"
	"time"
)

func main() {

	// not concurrent second line will wait first one to finish
	// repeat("Go")
	// repeat("Lang")

	// go keyword will make it concurrent
	// go keyword will create a goroutine
	// in that way we will be able to run thread at the same time.
	// go repeat("Go")
	// repeat("Lang")

	// if we add go keyword to both lines?
	// What did happen :D
	// we created two goroutine and
	// main function continue to run with next line
	// however because main routine reached end of its life
	// the created go routine will also die.
	// there will be no output.
	// go repeat("Go")
	// go repeat("Lang")

	// to fix this problem
	go repeat("Go")
	go repeat("Lang")

	// make main thread waits for goroutines work
	// we will have outputs in terminal for 2 seconds
	// time.Sleep(time.Second * 2)

	// another option
	// main function will wait io stream from user
	// if we press any key then main and other goroutines will die
	fmt.Scanln()

}

func repeat(word string) {

	for i := 0; true; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * 500)
	}

}
