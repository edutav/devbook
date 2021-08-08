package middlewares

import (
	"devbook/src/autenticacao"
	"devbook/src/respostas"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s - %s - %s", r.Method, r.Host, r.RequestURI)
		next(rw, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(rw, http.StatusUnauthorized, erro)
			return
		}
		next(rw, r)
	}
}
