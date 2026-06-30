package main

import (
	"fmt"
	"time"
)

func main() {
	var nome string
	fmt.Println("aaaaaaaa")
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Hello, World! 1")
		nome = "João"
		fmt.Println(nome)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Hello, World! 2")
		nome = "Maria"
		fmt.Println(nome)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello, World! 3")
		nome = "Pedro"
		fmt.Println(nome)
	}()

	fmt.Println("bbbbbbbb")

	select {}
} 