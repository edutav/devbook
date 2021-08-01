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
	statement, erro := ur.db.Prepare("insert into usuarios(nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
