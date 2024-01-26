package main

import "fmt"

// Declarar as variaveis na função diretamente e deixar o return limpo

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	soma, subtracao := calculosMatematicos(20, 10)
	fmt.Println(soma, subtracao)
}
