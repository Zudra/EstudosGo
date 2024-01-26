package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Array é uma lista de váriaveis
	// O array só permite 1 tipo
	// Um array vazio vai colocar o valor de espaços com o valor zero do tipo do seu array

	var array1 [5]string
	fmt.Println(array1)

	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Para não especificar a quantidade de posições

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Porém, não é dinâmico, não é possivel adicionar valores por fora do array
	// Já o slice têm tamanho dinâmico, mas ainda com a limitação do tipo

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// Slice não é um array

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// É possível adicionar um valor no Slice, ele adiciona um novo valor no seu slice com uma posição a mais

	slice = append(slice, 18)
	fmt.Println(slice)

	// Slice é um pedaço de um array
	// O primeiro index é inclusivo, o último é exclusivo
	// O slice aponta para um array, como um ponteiro

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// É póssivel alterar o valor de uma posição

	array2[1] = "Posição alterada"

	// Array interno

	fmt.Println("-----------------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacidade

	// O make criou um array de capacidade 15, e retorna um slice de 10 posições

	fmt.Println("-----------------------------")
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacidade

	// Agora ele ta na capacidade "máxima"
	fmt.Println("-----------------------------")
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacidade

	// Porém, ele pega o tamanho do array e cria um dobrando o seu valor de capacidade
	// Não é necessário especificar a capacidade do array no make

	fmt.Println("-----------------------------")
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// A capacidade do array vai ser deduzida pelo tamanho dele

	fmt.Println("-----------------------------")
	slice4 = append(slice4, 6)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
