package main

import "fmt"

// funções variáticas são usadas quando você precisa de vários parametros

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// É possivel combinar uma função com parametros fixos com parametros variáticos
// Porém, apenas pode ter um parametro variático

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	totalDaSoma := soma(1, 2, 5, 132, 56, 10, 44, 66)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo!", 01, 02, 03, 04, 05)

}
