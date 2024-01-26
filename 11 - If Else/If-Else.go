package main

import "fmt"

func main() {
	fmt.Println("If Else")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// Criar uma variável dentro do if causa essa variável ser usada apenas dentro deste if, não pode ser chamada fora

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número esta entre 0 e -10")
	}

}
