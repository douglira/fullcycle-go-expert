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

func StructsLesson(titleTemplate, title string) {
	fmt.Printf(titleTemplate, title)
	p1 := Customer{
		Name:   "Douglas",
		Age:    28,
		Enable: true,
	}
	p1.Address.Country = "Brazil"
	println(p1.Address.Country)
}
