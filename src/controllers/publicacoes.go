package controllers

import (
	"devbook/src/autenticacao"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// adiciona uma nova publicaçãono banco de dados
func CriarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	publicacao := models.Publicacao{}
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)
	publicacao.ID, erro = repositorioPublicacoes.Criar(publicacao)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusCreated, publicacao)
}

// traz as publicações que aparem no feed do usuário
func BuscarPublicacoes(rw http.ResponseWriter, r *http.Request) {

}

// traz uma unica publicação
func BuscarPublicacao(rw http.ResponseWriter, r *http.Request) {

}

// altera os dados de uma publicação
func AtualizarPublicacao(rw http.ResponseWriter, r *http.Request) {

}

// remove os dados de umaq publicação
func DeletarPublicacao(rw http.ResponseWriter, r *http.Request) {

}
