package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

// Testes automatizados são uma função que testará outra função para ver se essa
// função retornará o resultado esperado

func main() {
	tipoEndereco := enderecos.TiposDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
