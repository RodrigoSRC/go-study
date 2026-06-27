package main

import (
	"fmt"
	"os"
) 

func lerArquivo(nome string) (*os.File, error) {
	arq, err := os.Open(nome)
	if err != nil {
		return nil, err
	}
	defer arq.Close()
	return arq, nil
}

func main() { 
	fmt.Println("Hello, World!")
	_, err := lerArquivo("rotina.go")
	if err != nil {
		// fmt.Println("Erro ao ler o arquivo:", err)
		panic("Erro ao ler o arquivo: " + err.Error())	
	}

	// fmt.Println("Arquivo lido com sucesso")
}

func init() {
	fmt.Println("Init")
}