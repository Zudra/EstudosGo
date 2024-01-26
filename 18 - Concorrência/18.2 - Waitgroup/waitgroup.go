package main

import (
	"fmt"
	"sync"
	"time"
)

// Concorrência != paralelismo
//Paralelismo é quando duas tarefas estão sendo executadas simultaneamente
//Concorrência não está sendo executada necessariamente simultaneamente
//Concorrência executa uma tarefa sem a outra ter terminado, elas revezam o tempo de execução
//Mas elas também podem ser executadas simultaneamente caso seu processador tenha mais de um núcleo

func main() {

	// o método go colocado antes de invocar a função vai fazer que crie uma concorrência

	// Waitgroup é um método de sincronização das goroutines

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() { // Função anônima com goroutine
		escrever("Olá Mundo!")
		waitGroup.Done() // -1
	}()

	go func() { // Função anônima com goroutine
		escrever("Programando em GO!")
		waitGroup.Done() // -1
	}()

	go func() { // Função anônima com goroutine
		escrever("Goroutine 3!")
		waitGroup.Done() // -1
	}()

	go func() { // Função anônima com goroutine
		escrever("Goroutine 4!")
		waitGroup.Done() // -1
	}()

	// o go antes das func anonima, vai fazer ela não chegar no wait group antes do loop terminar

	waitGroup.Wait() // Vai dizer para a função esperar o waitgroup chegar em 0

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
