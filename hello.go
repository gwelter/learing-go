package main

import "fmt"

func main() {
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Inicial monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)
}
