package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func seno(s []string) string {
	co, _ := strconv.ParseFloat(s[0], 64)
	hi, _ := strconv.ParseFloat(s[2], 64)

	seno := co / hi
	senoConvertido := strconv.FormatFloat(seno, 'f', -1, 64)
	return "Seno: " + senoConvertido

}

func coseno(s []string) string {
	ca, _ := strconv.ParseFloat(s[1], 64)
	hi, _ := strconv.ParseFloat(s[2], 64)

	coseno := ca / hi
	cosenoConvertido := strconv.FormatFloat(coseno, 'f', -1, 64)
	return "\nCoseno: " + cosenoConvertido

}

func tangente(s []string) string {
	co, _ := strconv.ParseFloat(s[0], 64)
	ca, _ := strconv.ParseFloat(s[1], 64)

	tangente := co / ca
	tangenteConvertida := strconv.FormatFloat(tangente, 'f', -1, 64)
	return "\nTangente: " + tangenteConvertida

}

func main() {

	fmt.Println("Servidor aguardando conexões...")

	var wg sync.WaitGroup
	// chSeno := make(chan string)
	// chCoseno := make(chan string)
	// chTangente := make(chan string)

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

		valoresDoTriangulo := strings.Split(mensagem, " ")

		fmt.Print("Mensagem recebida:", valoresDoTriangulo)

		wg.Add(3)
		go func(wg *sync.WaitGroup) string {
			defer wg.Done()

			// chSeno <- seno(valoresDoTriangulo)
			time.Sleep(1 * time.Second)
			fmt.Println("fim go routine 1...")
			return "chSeno"
		}(&wg)

		go func(wg *sync.WaitGroup) string {
			defer wg.Done()
			// chTangente <- tangente(valoresDoTriangulo)
			fmt.Println("fim go routine 3...")

			return "chTangente"

		}(&wg)

		go func(wg *sync.WaitGroup) string {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			// fmt.Println(<-chSeno)

			// chCoseno <- coseno(valoresDoTriangulo)

			fmt.Println("fim go routine 2...")
			return "chCoseno"

		}(&wg)
		fmt.Println("Aguardando...")

		wg.Wait()

		fmt.Println("Fim...")

		// s := <-chSeno
		// sn := <-chCoseno
		// t := <-chTangente

		// close(chSeno)
		// close(chCoseno)
		// close(chTangente)

		// conexao.Write([]byte(string(s + sn + t + "\r\n")))

		conexao.Write([]byte(string(mensagem + "\r\n")))

	}
}
