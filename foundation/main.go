package main

import "fmt"

type ID int

var a ID = 1

func main() {
	/*
		%T - Mostra o tipo da variável
	*/
	fmt.Printf("O tipo de ID é %T \n", a)
	// R: O tipo de ID é main.ID

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("1- len=%d cap=%d %v\n", len(s), cap(s), s)
	/*
		Remove os elementos desde o índice informado até o último índice. Neste
		exemplo, estamos removendo do índice "2" a diante. O Slice retornado irá
		conter a quantidade de elementos informados (2 elementos) desde o início
		do Slice.

		Portanto, o ":" corta e retorna o Slice até o índice 2, sendo este índice
		exclusivo do retorno -> [10 20]
	*/
	fmt.Printf("2- len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	// R: "2- len=2 cap=10 [10 20]"
	fmt.Printf("3- len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	// R: "3- len=0 cap=10 []"
	/*
		Ignora a quantidade de elementos informados (neste exemplo 3 elementos)
		desde o ínicio do Slice e retorna o restante. Desta maneira, a capacidade
		do Slice é reduzida com o valor da quantidade de itens sobrados.

		Portanto, o ":" corta e retorna o Slice do índice 3 até o final, sendo este
		índice inclusivo -> [40 50 60 70 80 90 100]
	*/
	fmt.Printf("4- len=%d cap=%d %v\n", len(s[3:]), cap(s[3:]), s[3:])
	// R: "4- len=7 cap=7 [40 50 60 70 80 90 100]"

	/*
		Quando um elemento é adicionado a um Slice que não possui mais capacidade
		disponível, o Go dobra a capacidade deste Slice.
	*/
	s = append(s, 110)
	fmt.Printf("5- len=%d cap=%d %v\n", len(s), cap(s), s)
	// R: "5- len=11 cap=20 [10 20 30 40 50 60 70 80 90 100 110]"
}
