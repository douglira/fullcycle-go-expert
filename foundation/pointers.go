package foundation

import (
	"strings"
)

type Car struct {
	Brand        string
	BrandPointer *string
}

func PointersLesson() {
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
	c1 := Car{Brand: "bmw"}
	toUppercase(c1)
	println(c1.Brand)
	// R: bmw
	toUppercasePointer(&c1)
	println(c1.Brand)
	// R: BMW

	c2 := Car{Brand: "audi"}
	c2.toUppercase()
	println(c2.Brand)
	// R: audi

	audiBrand := "audi"
	c3 := Car{BrandPointer: &audiBrand}
	c3.toUppercasePointer()
	println(*c3.BrandPointer)
	// R: AUDI

	c4 := &Car{Brand: "porsche"}
	c4.toUppercasePointer2()
	println(c4.Brand)
	// R: PORSCHE
}

func toUppercase(c Car) {
	c.Brand = strings.ToUpper(c.Brand)
}

func toUppercasePointer(c *Car) {
	c.Brand = strings.ToUpper(c.Brand)
}

func (c Car) toUppercase() {
	brandLowercase := c.Brand
	c.Brand = strings.ToUpper(brandLowercase)
}

/*
Propriedade como ponteiro altera o valor no endereço de memória (MUTÁVEL)
*/
func (c Car) toUppercasePointer() {
	*c.BrandPointer = strings.ToUpper(*c.BrandPointer)
}

/*
Objeto como ponteiro altera o valor no endereço de memória (MUTÁVEL)
*/
func (c *Car) toUppercasePointer2() {
	c.Brand = strings.ToUpper(c.Brand)
}
