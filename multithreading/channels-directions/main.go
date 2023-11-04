package main

import "fmt"

/*
Através desta assinatura no argumento da função, indicamos a variável
referente ao channel somente pode ser utilizada para enviar informações
ao canal, impossibilitando de consumi-las (Send-only)
*/
func send(name string, hello chan<- string) {
	hello <- name
}

/*
Através desta assinatura no argumento da função, indicamos a variável
referente ao channel somente pode ser utilizada para consumir informações
do canal, impossibilitando realizar envios no mesmo (Receive-only)
*/
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go send("Hello", hello)
	read(hello)
}
