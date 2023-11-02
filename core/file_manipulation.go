package core

import (
	"bufio"
	"fmt"
	"os"
)

func FileManipulationLesson() {

	/*
		Criando e escrevendo no arquivo
	*/
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Arquivo escrito com sucesso"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("O arquivo possui %d bytes\n", size)

	/*
		Leitura de todo o arquivo
	*/
	fileRead, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileRead))

	/*
		Leituro do arquivo em partes (chunks)
	*/
	fileOpened, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(fileOpened)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
