package main

import (
	"fmt"
	"time"
)

// Concorrência != paralelismo
//Paralelismo é quando duas tarefas estão sendo executadas simultaneamente
//Concorrência não está sendo executada necessariamente simultaneamente
//Concorrência executa uma tarefa sem a outra ter terminado, elas revezam o tempo de execução
//Mas elas também podem ser executadas simultaneamente caso seu processador tenha mais de um núcleo

func main() {

	// o método go colocado antes de invocar a função vai fazer que crie uma concorrência

	go escrever("Olá Mundo!") // Goroutine
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
