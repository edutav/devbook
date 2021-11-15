package main

import (
	"devbook/src/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Println("VERSÃO:", config.Version)
}

func main() {
	config.Carregar()

	router := router.Gerar()

	fmt.Printf("\nExecutando API na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
