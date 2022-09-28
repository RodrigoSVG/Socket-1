package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

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
		//Moca a entrada do usuário
		valoresDoTriangulo := strings.Split(mensagem, " ")

		// escreve no terminal a mensagem recebida
		fmt.Print("Mensagem recebida:", valoresDoTriangulo)

		// operacao, _ := strconv.Atoi("3")

		// print(operacao, mensagem, err)

		// para um exemplo simples de processamento, converte a mensagem recebida para caixa alta
		// novamensagem := strings.ToUpper(mensagem)
		// outramensagem = 1 + 2
		// t := outramensagem
		go func() int {
			operacao, _ := strconv.Atoi(string(mensagem))
			d := 1 + operacao + 7
			return d
		}()

		// resultado := strconv.Itoa(operacao)

		// envia a mensagem processada de volta ao cliente
		teste := valoresDoTriangulo[1]
		vasdy := string(teste)
		conexao.Write([]byte(vasdy + "\n"))
		// conexao.Write([]byte(valoresDoTriangulo[1] + "\n"))
		// conexao.Write([]byte(valoresDoTriangulo[2] + "\n"))
	}
}
