package main

import "fmt"

// Usado para colocar várias informações sobre uma mesma coisa

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	// Struct sem valor, o seu valor zero é os campos do struct e jogar seus valores zero

	var u usuario
	fmt.Println(u)

	u.nome = "Murilo"
	u.idade = 18
	fmt.Println(u)

	// Pode se fazer o input dos campos do struct desta forma
	// É possível colocar structs dentro de structs

	enderecoexemplo := endereco{"Rua dos bobos", 0}
	usuario2 := usuario{"Murilo", 18, enderecoexemplo}
	fmt.Println(usuario2)

	// Se não tiver informação de outro campo do struct se faz assim:

	usuario3 := usuario{nome: "Murilo"}
	fmt.Println(usuario3)

}
