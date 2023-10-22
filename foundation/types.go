package foundation

import "fmt"

type ID int

var a ID = 1

func TypesLesson() {
	titleLesson := `
	########################################################################
	############################ TYPES ####################################
	########################################################################
	`
	fmt.Println(titleLesson)
	/*
		%T - Mostra o tipo da variável
	*/
	fmt.Printf("O tipo de ID é %T \n", a)
	// R: O tipo de ID é main.ID
}
