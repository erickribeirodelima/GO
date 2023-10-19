package main

import (
	"errors"
	"fmt"
)

func main() {
	// Números inteiros -- Existem 4 tipos de números inteiros no go
	// int8, int16, int32, int64

	var numero1 int = 28
	fmt.Println(numero1)

	numero2 := 100000000
	fmt.Println(numero2)

	// uint = não aceitar sinal de menos
	// Exemplo: var numero3 uint = -10000000

	// --------------------------------------------------

	// Números Reais = float32, float64
	// Só existem esses dois floats acima
	// Não pode usar simplesmente float, tem que especificar entre 32 ou 64

	var numero4 float32 = 10.32
	var numero5 float64 = 1000.578

	fmt.Println(numero4)
	fmt.Println(numero5)

	// ----------------------------------------------------

	// Strings

	var str string = "Teste"
	fmt.Println(str)

	// Valor Zero = quando você inicia uma variável sem valor
	// Exemplo:

	var texto string
	fmt.Println(texto)

	// -----------------------------------------------------

	// Booleano = bool
	// Famoso true ou false

	var booleano1 bool = true
	var booleano2 bool
	fmt.Println(booleano1) // true
	fmt.Println(booleano2) // false = pois está sem declarar

	// ------------------------------------------------------

	// Erro

	var erro error
	fmt.Println(erro) // nil = pois está com valor zero.

	var erro2 error = errors.New("Esse é o jeito certo de fazer um erro!")
	fmt.Println(erro2)

}
