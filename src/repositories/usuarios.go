package repositories

import (
	"database/sql"
	"devbook/src/models"
)

type UsuarioRepository struct {
	db *sql.DB
}

//Cria um repositotório de usuários
func NovoRepositorioUsuario(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db}
}

// Insise um usuário no banco
func (ur UsuarioRepository) Criar(usuario models.Usuario) (uint64, error) {
	return 0, nil
}
