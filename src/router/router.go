package router

import (
	"devbook/src/router/rotas"

	"github.com/gorilla/mux"
)

// vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
