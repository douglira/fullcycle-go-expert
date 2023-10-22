package foundation

import "fmt"

type Address struct {
	Street   string
	City     string
	Province string
	Number   string
	Country  string
}

type Customer struct {
	Name    string
	Age     int
	Enable  bool
	Address // Composição de Structs sem declarar explicitamente uma propriedade
}

func StructsLesson() {
	titleLesson := `
	########################################################################
	############################ STRUCTS ####################################
	########################################################################
	`
	fmt.Println(titleLesson)
	p1 := Customer{
		Name:   "Douglas",
		Age:    28,
		Enable: true,
	}
	p1.Address.Country = "Brazil"
	fmt.Println(p1.Address.Country)
}
