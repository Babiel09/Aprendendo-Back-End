package main

import (
	"fmt" //FMT é o pacote básico do GOLang
	"io/ioutil"
	"net/http" //O HTTP é para eu fazer requisições WEB
	"os"       //O OS serve para eu fazer algo em relação ao sistema operacional do usuário
	"strconv"
	"time"
)

// Defino essas variáveis globalmente pois elas serão usadas muitas vezes:
var siteses = []string{""}
var site string

const medidaTempo int = 5

func main() {

	for {
		perguntarComando()
		comando := escanearComando()

		switch comando {
		case 1:
			monitorarSite()
		case 2:
			mostrarOsLogs()
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
	fmt.Println("2- Mostrar Logs localmente")
	fmt.Println("3- Monitoramento mais robusto")
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
	res, err := http.Get(site)
	pegarErro()

	fmt.Println("Você fez um monitoramento desse(s) site(s): *", siteses, "*")
	fucaoDeLog(site, true)
	fmt.Println("Todos os sites que foram monitorados estão no arquivo: 'log.txt'")
	if err != nil {
		fmt.Println("Ocorreu um erro, virifique a url do site:", err, res)
	}
}
func repetirMonitoramento() {
	monitorarSite()
	fmt.Println("Monitorando...")
	fmt.Println("Execute : 'CTRL + C' para interromper o script quando for necessário")
	for {
		time.Sleep(time.Duration(medidaTempo) * time.Minute)
		pegarErro()

	}

}

// Criando uma função para pegar um erro:
func pegarErro() {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro inesperado: ", err, " verifique as informações da sua url.")
	}
	if res.StatusCode == 200 {
		fmt.Println("O site *", site, "* está no ar")
	}
}
func fucaoDeLog(site string, online bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02-01-2006 15:04:05") + " " + site + " *site online: " + strconv.FormatBool(online) + "* " + "\n") //uso o "\n" para pular uma linha
	arquivo.Close()
}

// Criando a função para mostrar os logs:
func mostrarOsLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
