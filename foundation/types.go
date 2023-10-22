package foundation

import "fmt"

type ID int

var a ID = 1

func TypesLesson() {
	/*
		%T - Mostra o tipo da variável
	*/
	fmt.Printf("O tipo de ID é %T \n", a)
	// R: O tipo de ID é main.ID
}
