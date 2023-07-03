package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	filePath := os.Args[1]
	serverAddress := "127.0.0.1:8080"

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Erro ao obter informações do arquivo:", err)
		return
	}

	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	// Enviar nome do arquivo e tamanho para o servidor
	_, err = conn.Write([]byte(fileName + ":" + fmt.Sprint(fileSize)))
	if err != nil {
		fmt.Println("Erro ao enviar nome do arquivo e tamanho:", err)
		return
	}

	buffer := make([]byte, 1)
	var sentBytes int64
	fmt.Println("Enviando arquivo", os.Args[1], "...")

	for {
		// Ler o arquivo em pedaços e enviar para o servidor
		_, err = file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}

		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println("Erro ao enviar o arquivo:", err)
			return
			sentBytes += int64(len(buffer))
			fmt.Printf("Progresso: %.2f%%\n", float64(sentBytes)/float64(fileSize)*100)
		}
	}

	fmt.Println("Arquivo enviado com sucesso!")
}
