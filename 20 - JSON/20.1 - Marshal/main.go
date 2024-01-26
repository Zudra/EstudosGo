package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	// Função marshal retorna o struct em JSON (em bytes) e um erro

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil { // Tratar o erro
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	// Para vermos em JSON usa a biblioteca bytes

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	// É possível usar um map em vez de um struct para construir um JSON

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
