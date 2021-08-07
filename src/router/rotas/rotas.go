package rotas

import (
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
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
		fmt.Println("Rota detectada ->", rota.Uri)
	}

	return r
}
