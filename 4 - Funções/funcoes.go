package main

import (
	"fmt"
)

// A função é um bloco de código reutilizável, você especifica a sua utilidade e chama o mesmo após com parametros

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// As funções podem ter mais de um retorno, é possivel colocar todos os tipos de retornos também

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// É possível usar a função como um tipo para ser colocada em uma variável

	var exemplo = func() {
		fmt.Println("Função F")
		return
	}

	exemplo()

	// Também é possivel passar parametros para essa função dentro da variável

	var exemplo2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := exemplo2("Texto da função 1")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	// Quando quiser apenas um resultado da função, e ignorar o outro, coloque um _

	resultadosSoma1, _ := calculosMatematicos(20, 10)
	fmt.Println(resultadosSoma1)

	// Ou

	_, resultadoSubtracao1 := calculosMatematicos(20, 10)
	fmt.Println(resultadoSubtracao1)
}
