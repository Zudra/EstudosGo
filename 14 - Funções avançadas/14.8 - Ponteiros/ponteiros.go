package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	// Sem o ponteiro, a variavel número mantem inalterada depois da execução da função

	fmt.Println("Ponteiros em funções")
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	fmt.Println("----------------------------")

	// Com o ponteiro, a alteração feita pela função vai ser feita no endereço de memória dessa variável

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
