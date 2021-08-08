package main

import (
	"devbook/src/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	//fmt.Println("URL do banco:", config.StringConexaoBanco)

	router := router.Gerar()

	fmt.Printf("\n Executando API na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
