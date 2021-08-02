package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta em Json
func JSON(rw http.ResponseWriter, statusCode int, dados interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if erro := json.NewEncoder(rw).Encode(dados); erro != nil {
		log.Fatal(erro)
	}

}

// Retorna uma mensagem em formato json
func Erro(rw http.ResponseWriter, statusCode int, erro error) {
	JSON(rw, statusCode, struct {
		Erro string `json:"erro"`
	}{Erro: erro.Error()})
}
