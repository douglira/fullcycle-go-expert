package foundation

// EXAMPLE 1
func Sum[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// EXAMPLE 2
type Number interface {
	/*
		Com o sinal "~" o Go permite que seja utilizado variáveis que possuam
		um tipo customizado, mas que por baixo dos panos, utilizam o mesmo
		valor primitivo.
		Neste caso da Generic, o tipo "Number" também aceita receber um valor
		do tipo "MyNumber" por utilizar um "int" por baixo dos panos.
	*/
	~int | float32 | float64
}

func Multiply[T Number | float64](m map[string]T) T {
	var sum T = 1
	for _, v := range m {
		sum *= v
	}
	return sum
}

type MyNumber int

/*
A keyword "comparable" permite com que seja feita a comparação de
igualdade entre os 2 elementos
*/
func Compare[T comparable](a T, b T) bool {
	return a == b
}

func GenericsLesson() {
	m := map[string]int{"Douglas": 1500, "Roberto": 3000}
	m1 := map[string]float64{"Douglas": 2550.25, "Roberto": 3000.10}
	m2 := map[string]MyNumber{"Douglas": 2550, "Roberto": 3000}
	println(Sum(m))
	println(Sum(m1))
	println(Multiply(m))
	println(Multiply(m1))
	println(Multiply(m2))

	println(Compare(10, 10))
	// R: true
	/*
		O comando abaixo não roda, pois ele espera que os 2 elementos sejam
		do mesmo tipo
	*/
	// println(Compare(10, 10.0))
	// R: default type float64 of 10.0 does not match inferred type int for T
}
