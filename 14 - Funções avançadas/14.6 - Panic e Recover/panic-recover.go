package main

import "fmt"

func recuperarExecucao() {
	// O recover pode ser usado em uma função que não está entrando em pânico, vai ser ignorada
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// Panic mata a execução do programa e todo o resto do programa não vai ser executado depois do panic
	// Panic não é um erro

	panic("A MÉDIA É EXATAMENTE 6!")

}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução!")
}
