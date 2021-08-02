package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Cria usuário no banco de dados
func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnprocessableEntity, erro)
		return
	}

	user := models.Usuario{}

	if erro = json.Unmarshal(corpoRequest, &user); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Preparar(); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositoryUsuario := repositories.NovoRepositorioUsuario(db)
	user.ID, erro = repositoryUsuario.Criar(user)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(rw, http.StatusCreated, user)
}

// Lista usuários no banco de dados
func ListarUsuarios(rw http.ResponseWriter, r *http.Request) {
	nomeNick := strings.ToLower((r.URL.Query().Get("usuario")))

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoryUsuario := repositories.NovoRepositorioUsuario(db)
	users, erro := repositoryUsuario.Buscar(nomeNick)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, users)

}

// Busca usuário no banco de dados
func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando usuário"))
}

// Atualiza usuário no banco de dados
func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Atualizando usuário"))
}

// Deleta usuário no banco de dados
func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Deletando usuário"))
}
