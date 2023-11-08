package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	dataChan := make(chan int)
	workersQuantity := 10000

	// init workers
	for i := 0; i < workersQuantity; i++ {
		go worker(i, dataChan)
	}

	// send data to be processed from Workers
	for i := 0; i < 10000000; i++ {
		dataChan <- i
	}
}
