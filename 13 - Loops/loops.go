package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	// No Go existe so o loop for, que repete o bloco de código até que a sua condição seja verdadeira

	for i < 10 {
		i++
		fmt.Println("Incrementando i...")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	// Também é possivel escrever a variável para apenas este escopo e não para fora

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	//clausula range, iterar um array ou um slice

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Para ter só os nomes e nao o index

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// iterar em um slice com for

	nomes1 := []string{"Erick", "Murilo", "Vitor"}

	for _, nome := range nomes1 {
		fmt.Println(nome)
	}

	// iterar uma string, ele retornara o valor da letra da tabela ASCII

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	// iterar um map, não é possivel fazer um for loop em um struct

	usuario := map[string]string{
		"nome":      "Murilo",
		"sobrenome": "Wolff",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Fazer um for infinito é dessa forma

	for {
		fmt.Println("Executando Infinitamente")
		time.Sleep(time.Second)
	}
}
