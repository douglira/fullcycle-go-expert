package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

/*
O WaitGroup irá aguardar uma quantidade de processos/operações a serem
finalizadas. Neste exemplo, serão 25 operações pois o número total de
loopings efetuados são 25.
*/
func WaitGroupLesson() {
	wg := sync.WaitGroup{}
	wg.Add(25)
	go task("A", &wg)
	go task("B", &wg)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			wg.Done()
		}
	}()
	wg.Wait()
}
