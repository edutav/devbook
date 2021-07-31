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

	fmt.Println("Executando API na porta", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
