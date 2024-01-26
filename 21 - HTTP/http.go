package main

import (
	"log"
	"net/http"
)

// funções das URis

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar páginas de usuários!"))
}

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz"))
}

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação web

	// hypertext transfer protocol

	// cliente - servidor

	// Request - Response

	// Rotas
	// URI - Indentificador do Recurso
	// Método - GET, POST, PUT, DELETE

	// ------------------------------------------------------------

	// Caso uma requisição venha nessa URI (/home), ela vai dar o retorno da função escrita

	http.HandleFunc("/", raiz)

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	// Cria uma rota na porta 5000, primeiro parametro

	log.Fatal(http.ListenAndServe(":5000", nil))
}
