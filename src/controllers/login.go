package controllers

import (
	"devbook/src/autenticacao"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/respostas"
	"devbook/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request) {
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

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoryUsuario := repositories.NovoRepositorioUsuario(db)
	userSaved, erro := repositoryUsuario.BuscarPorEmail(user.Email)
	if erro != nil {
		respostas.Erro(rw, http.StatusInternalServerError, erro)
		return
	}

	erro = seguranca.VerificarSenha(userSaved.Senha, user.Senha)

	if erro != nil {
		respostas.Erro(rw, http.StatusUnauthorized, erro)
		return
	}

	token, _ := autenticacao.CriarToken(userSaved.ID)
	fmt.Println(token)
	rw.Write([]byte(token))

	//respostas.JSON(rw, http.StatusOK, userSaved)
}
