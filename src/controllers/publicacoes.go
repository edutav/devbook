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

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["id_publicacao"], 10, 64)
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

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)
	publicacao, erro := repositorioPublicacoes.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusCreated, publicacao)
}

// traz uma unica publicação
func BuscarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)
	publicacoes, erro := repositorioPublicacoes.Buscar(usuarioID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusOK, publicacoes)

}

// altera os dados de uma publicação
func AtualizarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["id_publicacao"], 10, 64)
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

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)

	publicacaoSalvaNoBanco, erro := repositorioPublicacoes.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
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

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorioPublicacoes.Atualizar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// remove os dados de umaq publicação
func DeletarPublicacao(rw http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["id_publicacao"], 10, 64)
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

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)

	publicacaoSalvaNoBanco, erro := repositorioPublicacoes.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(rw, http.StatusForbidden, errors.New("não é possível deletar uma publicação que não seja sua"))
		return
	}

	if erro = repositorioPublicacoes.Deletar(publicacaoID); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// traz as publicações de um usuário especifico
func BuscarPublicacoesPorUsuario(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(params["usuario_id"], 10, 64)
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

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)
	publicacao, erro := repositorioPublicacoes.BuscarPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(rw, http.StatusCreated, publicacao)
}

// adiciona uma curtida na publicação
func CurtirPublicacao(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["id_publicacao"], 10, 64)
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

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)

	if erro = repositorioPublicacoes.Curtir(publicacaoID); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}

// remove uma curtida na publicação
func DescurtirPublicacao(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(params["id_publicacao"], 10, 64)
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

	repositorioPublicacoes := repositories.NovoRepositorioPublicacoes(db)

	if erro = repositorioPublicacoes.Descurtir(publicacaoID); erro != nil {
		respostas.Erro(rw, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(rw, http.StatusNoContent, nil)
}
