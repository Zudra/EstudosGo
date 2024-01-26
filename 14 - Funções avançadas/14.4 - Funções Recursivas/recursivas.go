package main

import "fmt"

// Funções recursivas chamam elas mesmas
// Depende que a sua execução funcione, depende de outra execução dela mesma
// Possui condição de parada

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(10)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(posicao))

}
