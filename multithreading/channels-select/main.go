package main

import "time"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 5)
		ch2 <- 2
	}()

	select {
	case msg1 := <-ch1:
		println("received ch1: ", msg1)
	case msg2 := <-ch2:
		println("received ch2: ", msg2)
	case <-time.After(time.Second * 4):
		println("channels timeout")

		/*
			Default case roda imediatamente e não aguarda os outros cases
		*/
		// default:
		// 	println("default")
	}
}
