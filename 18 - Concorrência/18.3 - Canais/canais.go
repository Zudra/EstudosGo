package main

import (
	"fmt"
	"time"
)

// Canais funcionam para enviar e receber dados

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	// O goroutine vai esperar até o canal receber o valor

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal
		// Para evitar o deadlock, se faz uma segunda variável bool para quebrar quando o canal for fechado
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// Também é possível fazer esse loop para repetir o canal desta maneira

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	// Fechará o canal após o loop ser finalizado

	close(canal)

}
