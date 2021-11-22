package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	// send 1, 2... 100 as n for jobs to
	// calculate nth fibonnnaci number
	for i := 0; i < 100; i++ {
		jobs <- i
	}

	// close channel when its done
	close(jobs)

	// we send 100 numbers to channel
	// we will receive 100 from the channel
	for j := 0; j < 100; j++ {
		fmt.Printf("fib %d.: %d\n", j, <-results)
	}

}

// jobs will only receive from channel
// results will only send message to channel
func worker(jobs <-chan int, results chan<- int) {

	// calculate nth fibonacci number and send to channel
	for n := range jobs {
		results <- fib(n)
	}

}

func fib(n int) int {

	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
