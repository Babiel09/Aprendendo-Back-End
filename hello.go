package main

import (
	"fmt"      //FMT é o pacote básico do GOLang
	"net/http" //O HTTP é para eu fazer requisições WEB
	"os"       //O OS serve para eu fazer algo em relação ao sistema operacional do usuário
	"time"
)

// Defino essas variáveis globalmente pois elas serão usadas muitas vezes:
var siteses = []string{""}
var site string

func main() {

	for {
		perguntarComando()
		comando := escanearComando()

		switch comando {
		case 1:
			monitorarSite()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 3:
			repetirMonitoramento()
		case 0:
			fmt.Println("Ok, até mais")
			os.Exit(0)

		default:
			fmt.Println("Não conheço este comando")
		}
	}
}

// Criando novas funções
// Função para perguntar o quê o usuário vai fazer
func perguntarComando() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Teste mais robusto")
	fmt.Println("0- Sair do Programa")
	return
}

// Função para escanear o comando:
func escanearComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando //Aqui é obrigatório eu retornar o comando sendo que eu declarei ele lá em cima
}

// Criando a função para monitorar o site:
func monitorarSite() {

	fmt.Println("Digite a URL do site para monitoramento:")

	//Pedir para o usuário escrever o site que deseja monitorar!
	fmt.Scan(&site)
	siteses = append(siteses, site)
	fmt.Println("Monitorando...")
	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		fmt.Println("O site *", site, "* está no ar")
	} else {
		fmt.Println("Deu um erro absurdo, verifique se o site está no ar ou se ele existe!")
		fmt.Println(res.StatusCode)
	}

	fmt.Println("Você fez um monitoramento desses sites: *", siteses, "*")

}
func repetirMonitoramento() {
	monitorarSite()
	for {
		time.Sleep(4 * time.Second)
		fmt.Println("Monitorando...")
		res, _ := http.Get(site)
		if res.StatusCode == 200 {
			fmt.Println("O site *", site, "* está no ar")
		} else {
			fmt.Println("Deu um erro absurdo, verifique se o site está no ar ou se ele existe!")
			fmt.Println(res.StatusCode)
		}
	}
}
