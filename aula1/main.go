package main

import (
	"fmt"
	"os"
) 

func main() {
	fmt.Println("Hello, World!")
	_, err := os.Open("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	// defer arq.Close()
	// arq, err = os.Open("arquivo.txt")

}

func init() {
	fmt.Println("Init")
}