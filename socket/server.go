package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	var receivedBytes int64
	var fileSize int64
	var fileName string

	// Receber nome do arquivo e tamanho do cliente
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao receber nome do arquivo e tamanho:", err)
		return
	}

	fileData := strings.Split(string(buffer), ":")
	fileName = fileData[0]
	fileSize = stringToInt64(fileData[1])

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Println("Recebendo arquivo...")

	for {
		// Receber o arquivo em pedaços e escrever no disco
		n, err := conn.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro ao receber o arquivo:", err)
			return
		}

		_, err = file.Write(buffer[:n])
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}

		receivedBytes += int64(n)
		fmt.Printf("Progresso: %.2f%%\n", float64(receivedBytes)/float64(fileSize)*100)

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Arquivo recebido com sucesso!")
}

func stringToInt64(str string) int64 {
	var result int64
	fmt.Sscanf(str, "%d", &result)
	return result
}

func main() {
	host := "0.0.0.0"
	port := "8080"

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor iniciado. Aguardando conexões...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar a conexão:", err)
			return
		}

		go handleClient(conn)
	}
}
