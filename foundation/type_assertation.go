package foundation

import "fmt"

func TypeAssertationLesson() {
	var name interface{} = "Douglas Lira"
	println(name)
	// R: (0x1097a60,0x10c7fc0)
	println(name.(string))
	// R: Douglas Lira
	result, ok := name.(int)
	fmt.Printf("O valor de 'result' é %v e o valor de 'ok' é %v\n", result, ok)
	// R: O valor de 'result' é 0 e o valor de 'ok' é false%
	/*
		Sem a verificação do segundo retorno, o Go lança um PANIC
	*/
	// result2 := name.(int)
	// println(result2)
	// R: panic: interface conversion: interface {} is string, not int
}
