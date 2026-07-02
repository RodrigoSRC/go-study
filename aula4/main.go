package main

import (
	cp "github.com/aoticombr/golang/component"
	lib "github.com/aoticombr/golang/lib"
)

func SortearIdade(ate int ) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ate)
}

func CriarArquivoGb(value int) {
	var arquivo cp.String
	arquivo.Delimiter = "\n"

	arquivo.Add("Nome | Idade | Dta Nasci | Logradouro | Bairro | Cidade")

	for i := 0; i < 10000000*value; i++ {
		Idade:= SortearIdade(100)
		Data:= time.Now()
		Data = Data.AddDate(0, 0, -Idade)
		linha:= fmt.Sprintf("Nome %d | %d | %s | Rua %d | Bairro %d | Cidade %d", i, Idade, Data.Format("02/01/2006"), i, i, i)
		arquivo.Add(linha)
	}

	err:= lib.ByteToSaveFile("aula4.txt", arquivo.Bytes())
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}
}

func main() {
	CriarArquivoGb(2)


	// fmt.Println("Arquivo salvo com sucesso!")
}