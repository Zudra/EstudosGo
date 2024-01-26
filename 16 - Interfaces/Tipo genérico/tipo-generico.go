package main

import "fmt"

// interface de tipo genérico aceita qualquer coisa, não tem restrição
// é uma interface vazia

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
