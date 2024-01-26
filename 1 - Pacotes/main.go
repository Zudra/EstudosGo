package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	var erro = checkmail.ValidateFormat("murilo.klug@gmail.com")
	fmt.Println(erro)

}
