package main

import "fmt"

// Switch é usado quando uma condição muito específica é necessária e pode ter várias condições específicas

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número invalido"
	}
}

// A condição pode ser escrita desta maneira também

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

// O return pode ser escrito desta maneira também

func diaDaSemana3(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
		// Esse argumento faz com que mesmo se o número do switch esteja nessa condição, fará ela cair direto
		// Na condição de baixo, sem avaliar o case
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	fmt.Println("-------------------")
	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)

	fmt.Println("-------------------")
	dia3 := diaDaSemana3(1)
	fmt.Println(dia3)
}
