package main

import (
	"errors"
	"fmt"
)

func main() {
	//Existem 4 tipos de int que suportam diferentes tamanhos de dados

	var numero int64 = 1000000000000
	fmt.Println(numero)

	// Quando nao for especificado o tipo de int, ele vai usar a arquitetura do seu
	// computador, se ele for 64 bits, ele vai usar o int64

	// uint: é um int que não aceita números negativos, portanto daria erro de
	// compilação caso fosse usado com número negativo

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// Existem alias/apelidos para o int
	// INT32 = RUNE

	var numero3 rune = 123456
	fmt.Println(numero3)

	// UINT8 = BYTE

	var numero4 byte = 123
	fmt.Println(numero4)

	// Números reais (float) apenas tem duas versões, a 64 e 32 bits, funciona com a omissão

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12321.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12378.128739
	fmt.Println(numeroReal3)

	// Não existe char no go, portanto o char se torna um número na tabela ASCII,
	// então se usa as aspas singulares e não duplas

	char := 'A'
	fmt.Println(char)

	// Todo tipo de dado tem seu valor zero, quando você não da um valor pra mesma
	// int vai dar um valor 0, string vai dar vazia

	var texto int
	fmt.Println(texto)

	// Booleano é o mesmo que em outras linguagens, pode ser omitido
	// também pode ser deixado sem valor, o seu número zero é False

	var booleanos bool
	fmt.Println(booleanos)

	// Go possui também o tipo de dado chamado "erro", o seu tratamento é diferente
	// Seu valor zero é nul (nulo)

	var erro error
	fmt.Println(erro)

	// error é o tipo de dado, errors é um pacote para escrever um erro interno

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

}
