package main

import (
	"fmt"
	"os"
)

func main() {
	perguntarComando()
	comando := escanearComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)

	default:
		fmt.Println("Não conheço este comando")
	}
}

// Criando novas funções
// Função para perguntar o quê o usuário vai fazer
func perguntarComando() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	return
}

// Função para escanear o comando:
func escanearComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando //Aqui é obrigatório eu retornar o comando sendo que eu declarei ele lá em cima
}
