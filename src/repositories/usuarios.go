package repositories

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
)

type usuarioRepository struct {
	db *sql.DB
}

//Cria um repositotório de usuários
func NovoRepositorioUsuario(db *sql.DB) *usuarioRepository {
	return &usuarioRepository{db}
}

// Insise um usuário no banco
func (ur usuarioRepository) Criar(usuario models.Usuario) (uint64, error) {
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
func (ur usuarioRepository) Buscar(nomeNick string) ([]models.Usuario, error) {
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

func (ur usuarioRepository) BuscarPorId(id uint64) (models.Usuario, error) {
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

func (ur usuarioRepository) Atualizar(id uint64, usuario models.Usuario) error {
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

func (ur usuarioRepository) Deletar(id uint64) error {
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

func (ur usuarioRepository) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := ur.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	user := models.Usuario{}

	if linha.Next() {
		if erro = linha.Scan(&user.ID, &user.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return user, nil
}

// Permite que um usuário siga o outro
func (ur usuarioRepository) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := ur.db.Prepare("insert ignore into seguidores(usuario_id, seguidor_id) values(?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// Permite que um usuário pare de siguir o outro
func (ur usuarioRepository) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := ur.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// Busca todos os seguidores de um usuário
func (ur usuarioRepository) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := ur.db.Query(`select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?`, usuarioID)
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

// Busca todos os usuários na base
func (ur usuarioRepository) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := ur.db.Query(`select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?`, usuarioID)
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

// buscar senha por ai na base
func (ur usuarioRepository) BuscarSenha(id uint64) (string, error) {
	linha, erro := ur.db.Query("select senha from usuarios where id = ?", id)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	user := models.Usuario{}

	if linha.Next() {
		if erro = linha.Scan(&user.Senha); erro != nil {
			return "", erro
		}
	}

	return user.Senha, nil
}

// Altera a senha de um usuário
func (ur usuarioRepository) AtualizarSenha(id uint64, senhaComHash string) error {
	statement, erro := ur.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senhaComHash, id); erro != nil {
		return erro
	}

	return nil
}
