package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	/*
		Fechando o canal para evitar um deadlock error, pois o channel
		sempre irá aguardar um valor ser enviado/consumido.
	*/
	close(ch)
}
