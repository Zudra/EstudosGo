package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Parecido com o struct, porém tem chaves para os valores, a chaves tem que ter os mesmos tipos
	// Os valores também

	usuario := map[string]string{
		"Nome":      "Murilo",
		"Sobrenome": "Wolff",
	}

	fmt.Println(usuario)

	// Para pegar uma chave específica apenas é assim

	fmt.Println(usuario["Nome"])

	// É possível aninhar maps

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Murilo",
			"ultimo":   "Wolff",
		},
		"curso": {
			"nome":   "Eng de Software",
			"campus": "Fag Bloco 4",
		},
	}

	fmt.Println(usuario2)

	// É possível deletar uma chave dentro do map com um comando nativo do GO

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	// Também adicionar uma chave nova

	usuario2["Signo"] = map[string]string{
		"nome": "Peixes",
	}

	fmt.Println(usuario2)
}
