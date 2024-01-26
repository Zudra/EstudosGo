package main

import "fmt"

func main() {
	// Dá para deixar o tipo da variável implicito ou explicito

	var variavel1 = "Variável 1"
	variavel2 := "Variável 2"

	// O go não deixa você colocar algo que nao vai ser usado. vai ser um erro de compilação

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Também é possivel fazer uma lista com as variáveis

	var (
		variavel3 string = "HEHE"
		variavel4 string = "HAHA"
	)

	fmt.Println(variavel3, variavel4)

	// Não é necessário especificar o tipo da variável
	// pode ser escrito dessa forma também :
	//

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// Constantes são valores que não podem ser alterados

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Para inverter os valores das variáveis pode ser feito desta maneira

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
