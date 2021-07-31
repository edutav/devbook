package controllers

import "net/http"

// Cria usuário no banco de dados
func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Criando usuário"))
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
