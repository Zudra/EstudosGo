package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++

	// Apenas a variavel1 vai mudar, pois a variavel2 está com uma cópia da variavel1 antes de ser acrescida em 1

	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória

	var variavel3 int
	var ponteiro *int

	variavel3 = 100

	// o & vai colocar o endereço de memória da váriavel colocada no ponteiro

	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)

	// Se quiser ver o valor que está atribuido no endereço de memória, se coloca o * antes do ponteiro

	variavel3 = 150

	// Mesmo que o valor da variavel que o ponteiro está apontando mude, o endereço de memória não irá mudar

	fmt.Println(variavel3, *ponteiro)
}
