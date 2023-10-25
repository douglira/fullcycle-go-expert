package core_packages

import "fmt"

func DeferLesson() {
	/*
		A ordem de execuão de múltiplas chamadas utilizando "defer" ocorrem
		de maneira Last In First Out (LIFO)
	*/
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
	defer fmt.Println("Quarta linha")
	// R: Segunda linha
	// R: Terceira linha
	// R: Quarta linha
	// R: Primeira linha
}
