package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const URL = "http://viacep.com.br/ws/%s/json/"

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf(URL, cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler body: %v\n", err)
		}
		end := ViaCEP{}
		if err := json.Unmarshal(res, &end); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter json: %v\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", end.Cep, end.Localidade, end.Uf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao registrar infos no arquivo: %v\n", err)
		}
	}
}
