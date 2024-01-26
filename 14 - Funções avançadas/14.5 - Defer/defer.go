package main

import "fmt"

// defer adia a execução de um pedaço de código até o último momento
// Ou antes, do return

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {

	fmt.Println("Defer")
	defer funcao1()
	// DEFER = Adiar
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))

}
