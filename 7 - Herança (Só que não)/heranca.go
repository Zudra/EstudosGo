package main

import "fmt"

// Herança é para colocar informações de outro struct que tem os mesmos em outro
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 180}
	fmt.Println(p1)

	// é possível usar a struct de cima dentro de outra, desde que ela esteja descrita na outra e ja tenha valores

	e1 := estudante{p1, "Engenharia de Software", "Fag"}
	fmt.Println(e1)

	// também é possivel pegar uma declaração especifica de dentro da struct

	fmt.Println(e1.altura)
}
