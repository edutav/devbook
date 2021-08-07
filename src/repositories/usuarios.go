package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
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

// Buscar um usuário no banco baseado no filtro de nome ou nick
func (ur UsuarioRepository) Buscar(nomeNick string) ([]models.Usuario, error) {
	nomeNick = fmt.Sprintf("%%%s%%", nomeNick)

	linhas, erro := ur.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?", nomeNick, nomeNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []models.Usuario

	for linhas.Next() {
		var user = models.Usuario{}

		if erro = linhas.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur UsuarioRepository) BuscarPorId(id uint64) (models.Usuario, error) {
	linha, erro := ur.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", id)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	user := models.Usuario{}

	if linha.Next() {
		if erro = linha.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return user, nil
}

func (ur UsuarioRepository) Atualizar(id uint64, usuario models.Usuario) error {
	statement, erro := ur.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (ur UsuarioRepository) Deletar(id uint64) error {
	statement, erro := ur.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}
