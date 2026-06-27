package main

import (
	"fmt"
	"os"
) 

const msg_erro = "deu um [%s] %s %d xxxxxxx"

func lerArquivo(nome string) (*os.File, error) {
	arq, err := os.Open(nome)
	if err != nil {
		return nil, err
	}
	return arq, nil
}

func main() { 
	fmt.Println("Hello, World!")
	_, err := lerArquivo("rotina.go")
	if err != nil {
		fmt.Println(err.Error())
		err = fmt.Errorf(msg_erro, err.Error(), "aconteceu nao funcao ler arquivo", 404)
		fmt.Println(err)
		err = fmt.Errorf(msg_erro, err.Error(), "hhhhh", 500)
		fmt.Println(err)
		panic("deu um erro muito grave, erro:" + err.Error())
	}

	fmt.Println("Arquivo lido com sucesso")
}

func init() {
	fmt.Println("Init")
}