package main

import (
	cp "github.com/aoticombr/golang/component"
	lib "github.com/aoticombr/golang/lib"
)

func CriarArquivoGb(value int) {
	var arquivo cp.String

	arquivo.Add("Nome | Idade | Dta Nasci | Logradouro | Bairro | Cidade")

	err:= lib.ByteToSaveFile("aula4.txt", arquivo.Bytes())
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}
}

func main() {


	// fmt.Println("Arquivo salvo com sucesso!")
}