package main

import (
	"fmt"
)

func main() {
	fmt.Println("1- Inicar o monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
}
