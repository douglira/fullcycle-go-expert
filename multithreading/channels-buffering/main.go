package main

func main() {
	/*
		Configurando o buffer do channel, basicamente a quantidade limiite de dados
		que será trafegada pelo channel. Desta maneira, não é necessário aguardar
		o dado ser consumido no channel para enviar um outro dado, fazendo com
		que seja possível enviar (acumular) os dados no channel para serem lidos
		posteriormente. Importante tomar cuidado com a quantidade de buffer e o
		tamanho dos dados a serem "bufferizados", pois isto pode impactar
		negativamente no consumo de memória
	*/
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
