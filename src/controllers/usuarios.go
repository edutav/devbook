package controllers

import (
	"devbook/src/autenticacao"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if erro = user.Preparar("cadastro"); erro != nil {
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
	params := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	user, erro := repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	if user.ID == 0 {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, user)

}

// Atualiza usuário no banco de dados
func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	usuarioIDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	if usuarioIDToken != usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New(
			"não é possivel atualizar um usuário que não seja o seu"),
		)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	usuario := models.Usuario{}
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edição"); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// Deleta usuário no banco de dados
func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	usuarioIDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	if usuarioIDToken != usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possivel deletar um usuário que não seja o seu"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// Permite que um usuário siga o outro
func SeguirUsuario(rw http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// Permite que um usuário pare de siguir outro
func PararDeSeguirUsuario(rw http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível parar de seguir você mesmo"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// Busca seguidores de um usuário
func BuscarSeguidores(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioUsuario(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, seguidores)

}
