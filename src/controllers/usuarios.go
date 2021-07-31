package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Cria usuário no banco de dados
func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Criando usuário"))
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	user := models.Usuario{}

	if erro = json.Unmarshal(corpoRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositoryUsuario := repositories.NovoRepositorioUsuario(db)
	repositoryUsuario.Criar(user)
}

// Lista usuários no banco de dados
func ListarUsuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Listando usuários"))
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
