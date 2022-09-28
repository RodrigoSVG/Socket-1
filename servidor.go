package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func areaCirculo(cliente string) string {
	fmt.Println(cliente)
	// var r float64 = float64(cliente)
	// fmt.Println(r)
	// p := 3.14

	// A := p * (r * r)

	// teste := strconv.FormatFloat(A, 'f', -1, 64)

	return "teste"
}

func main() {

	fmt.Println("Servidor aguardando conexões...")

	// ouvindo na porta 8081 via protocolo tcp/ip
	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
		/* Neste nosso exemplo vamos convencionar que a saída 3 está reservada para erros de conexão.
		IMPORTANTE: defers não serão executados quando utilizamos os.Exit() e a saída será imediata */
		os.Exit(3)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()

	fmt.Println("Conexão aceita...")
	// rodando loop contínuo (até que ctrl-c seja acionado)
	for {
		// Assim que receber o controle de nova linha (\n), processa a mensagem recebida
		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}
		hi := strings.Split(mensagem, "\b")
		fmt.Println(hi[0])
		// mensagem2 := mensagem
		// novaMensagem, _ := strconv.ParseFloat(mensagem, 64)
		//Moca a entrada do usuário
		// valoresDoTriangulo := strings.Split(mensagem, " ")

		// escreve no terminal a mensagem recebida
		fmt.Print("Mensagem recebida:", mensagem)

		// envia a mensagem processada de volta ao cliente

		texto := string(mensagem)
		fmt.Println(texto)
		teste, err := strconv.Atoi(mensagem)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(teste)

		testando := areaCirculo(texto)

		// outramensagem := " teste" + mensagem
		// r := 10.0
		// p := 3.14

		// A := p * (r * r)

		// teste := strconv.FormatFloat(A, 'f', -1, 64)

		conexao.Write([]byte(testando + "\n"))

	}
}
