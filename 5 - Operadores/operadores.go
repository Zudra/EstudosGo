package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Não é possivel fazer qualquer operação com duas variáveis que tenham 2 tipos diferentes, por exemplo
	// Usar int16 com int32

	var num2 int16 = 30
	var num3 int16 = 25
	soma2 := num2 + num3
	fmt.Println(soma2)

	// fim dos aritméticos

	// Atribuição

	var variavel string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel, variavel2)

	// fim dos operadores de atribuição

	// Operadores relacionais
	// Sempre retornam True ou False

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// fim dos relacionais

	// Operadores lógicos

	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// fim dos operadores lógicos

	// Operadores unários

	fmt.Println("--------------------")
	numero := 10
	numero++     // numero =  numero + 1
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--    // numero = numero - 1
	numero -= 4 // numero = numero - 20

	numero *= 3 // numero = numero * 3
	numero /= 3 // numero = numero / 3
	numero %= 3 // numero = numero % 3

	fmt.Println(numero)

	// Fim dos operadores unários

	// Operadores
}
