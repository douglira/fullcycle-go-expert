package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	/*
		Adicionando 10 operações referentes às 10 iterações do
		for loop no publish()
	*/
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)

	/*
		WaitGroup aguardando o processamento das goroutines
	*/
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
