package main

import (
	"fmt"

	"linguagem-go/calc"
)

// Declarando variável
var nome string

func main() {
	// Atribuindo valor a variável
	nome = "Erick"

	// Declarando e atribuindo valor a variável
	// O sinal ':=' faz declaração e atribuição (só usa na hora de criar a variável)
	// O sinal '=' faz só atribuição
	idade := 20
	altura := 1.78
	formado := false
	descricao := `linguagem
	
	go
	`

	// Printando as informações
	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(altura)
	fmt.Println(formado)
	fmt.Println(descricao)

	// Formando o print
	fmt.Printf("%v \n", nome)
	fmt.Printf("%v \n", idade)
	fmt.Printf("%v \n", altura)
	fmt.Printf("%v \n", formado)
	fmt.Printf("%v \n", descricao)

	// Verificiando o tipo das variáveis
	fmt.Printf("%T \n", nome)
	fmt.Printf("%T \n", idade)
	fmt.Printf("%T \n", altura)
	fmt.Printf("%T \n", formado)
	fmt.Printf("%T \n", descricao)

	resultado := calc.Plus(1, 1)
	resultado_2 := calc.Plus10(10)
	fmt.Printf("%v \n", resultado)
	fmt.Printf("%v \n", resultado_2)
	fmt.Println(calc.Show)
}
