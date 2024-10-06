package main

import "reflect"

func main() {
	println("Mascote fofo demais")
	var nome string = "Teste"
	println(nome, reflect.TypeOf(nome))
}
