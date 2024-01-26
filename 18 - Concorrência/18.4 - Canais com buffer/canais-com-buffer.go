package main

import "fmt"

func main() {

	// Buffer especifica a capacidade do canal, no caso o 2 abaixo
	// Só bloqueia quando atingir a capacidade máxima

	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	// se eu escrevesse mais um canal, geraria um deadlock

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
