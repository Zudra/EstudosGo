package main

import "fmt"

// Método dá ações a alguma estrutura
// Sempre está grudado nessa estrutura, diferente das funções
// a sintaxe é a mesma da função porem com o parenteses a mais para anexar a estrutura usada

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Para alterar a estrutura, se passa o ponteiro

func (u *usuario) fazerAniversario() {
	u.idade++
}

type usuario struct {
	nome  string
	idade uint8
}

func main() {

	fmt.Println("Métodos")

	usuario1 := usuario{"Murilo", 18}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Erick", 20}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	println(maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
