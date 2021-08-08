package rotas

import (
	"devbook/src/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Representa todas as rotas da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Coloca todas as rotas dentro no Router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		fmt.Printf("\n Rota detectada -> %s : Método -> %s", rota.Uri, rota.Metodo)
		if rota.RequerAutenticacao {
			r.HandleFunc(
				rota.Uri,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
