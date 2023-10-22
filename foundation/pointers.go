package foundation

import (
	"fmt"
	"strings"
)

type Car struct {
	brand string
}

func PointersLesson(titleTemplate, title string) {
	fmt.Printf(titleTemplate, title)

	/*
		Declara uma nova variável com o mesmo ponteiro da variável
		que está sendo atribuída
	*/
	a := 10
	var pointer *int = &a
	println(pointer)
	// R: 0xc000092f68

	/*
		O valor da variável "a" é alterado para 20 também, pois
		ambas as variáveis compartilham o mesmo endereço de memória
	*/
	*pointer = 20
	println(a)
	// R: 20
	b := &a
	println(*b)
	// R: 20
	*b = 30
	println(a)
	// R: 30

	/*
		Para trabalhar com valores mutáveis, utilize ponteiros nos argumentos
		da função.
	*/
	c1 := Car{brand: "bmw"}
	toUppercase(c1)
	println(c1.brand)
	// R: bmw
	toUppercasePointer(&c1)
	println(c1.brand)
	// R: BMW
}

func toUppercase(c Car) {
	c.brand = strings.ToUpper(c.brand)
}

func toUppercasePointer(c *Car) {
	c.brand = strings.ToUpper(c.brand)
}
