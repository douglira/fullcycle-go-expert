package main

import (
	htmlTemplate "html/template"
	"os"
	"strings"
	textTemplate "text/template"
)

type Cursos []Curso

type Curso struct {
	Nome         string
	CargaHoraria int32
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func TemplateHTML() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	/*
		No "template.New()" é importante informar o nome do arquivo como base
	*/
	t := htmlTemplate.New("content.html")
	t.Funcs(htmlTemplate.FuncMap{"ToUpper": ToUpper})
	t = htmlTemplate.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go Expert", 40},
		{"Node Expert", 60},
		{"Python Expert", 20},
	})
	if err != nil {
		panic(err)
	}
}

func TemplateSimplesExemplo() {
	curso := Curso{"Go Expert", 40}
	t := textTemplate.Must(textTemplate.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	/*
		Utilizando Templates, podemos indicar o writer na função "Execute()" para
		realizar o parse dos valores. Exemplos de uso: formatar e-mails, logs, etc.
	*/
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}

func main() {
	// TemplateSimplesExemplo()
	TemplateHTML()
}
