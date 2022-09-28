package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// conectando na porta 8081 via protocolo tcp/ip na máquina local
	conexao, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	for {
		// lendo entrada do terminal
		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("Entre com o raio de um círculo")
		raio, err := leitor.ReadString('\n')

		// fmt.Print("Entre com os valores do triângulo (Cateto oposto, adjacente e hipotetusa): ")
		// co, err := leitor.ReadString(' ')
		// ca, err := leitor.ReadString(' ')
		// h, err := leitor.ReadString(' ')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		// escrevendo a mensagem na conexão (socket)
		fmt.Fprintf(conexao, raio+"\n")

		// ouvindo a resposta do servidor (eco)
		mensagem, err3 := bufio.NewReader(conexao).ReadString('\n')
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(3)
		}
		// escrevendo a resposta do servidor no terminal
		fmt.Println("Resposta do servidor: " + mensagem)
	}
}
