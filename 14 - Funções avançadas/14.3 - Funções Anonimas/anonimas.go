package main

import "fmt"

// uma função anônima é invocada de outra maneira, pois não tem nome, ela contém um parenteses no seu fim

func main() {

	func() {
		fmt.Println("Olá mundo")
	}()

	// Esse parenteses recebe o parametro recebido

	func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetro")

	// Se quiser que ela passe um retorno, precisa declarar uma variável para armazenar o retorno

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)

}
