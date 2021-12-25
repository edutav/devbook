package repositories

import (
	"database/sql"
	"devbook/src/models"
)

type publicacoesRepository struct {
	db *sql.DB
}

//Cria um repositotório de Publicações
func NovoRepositorioPublicacoes(db *sql.DB) *publicacoesRepository {
	return &publicacoesRepository{db}
}

// Insire uma publicação no banco
func (pr publicacoesRepository) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := pr.db.Prepare("insert into publicacoes(titulo, conteudo, autor_id) values(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Busca uma publicação no banco, por ID
func (pr publicacoesRepository) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := pr.db.Query(
		`
		select p.*, u.nick
		from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?
		`, publicacaoID,
	)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	publicacao := models.Publicacao{}

	if linha.Next() {
		if erro = linha.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Busca publicações no banco
func (pr publicacoesRepository) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := pr.db.Query(
		`
		select distinct p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id = ?
		order by 1 desc
		`, usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	publicacoes := []models.Publicacao{}

	for linhas.Next() {
		publicacao := models.Publicacao{}
		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// atualiza uma publicação no banco
func (pr publicacoesRepository) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := pr.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

// deleta uma publicação no banco
func (pr publicacoesRepository) Deletar(publicacaoID uint64) error {
	statement, erro := pr.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

// Busca publicações no banco de um determinado usuário
func (pr publicacoesRepository) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := pr.db.Query(
		`
		select distinct p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		where p.autor_id = ?
		order by 1 desc
		`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	publicacoes := []models.Publicacao{}

	for linhas.Next() {
		publicacao := models.Publicacao{}
		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// adiciona uma curtida na publicação
func (pr publicacoesRepository) Curtir(publicacaoID uint64) error {
	statement, erro := pr.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

// adiciona uma curtida na publicação
func (pr publicacoesRepository) Descurtir(publicacaoID uint64) error {
	statement, erro := pr.db.Prepare(`
		update publicacoes set curtidas = curtidas 
		CASE
			WHEN curtidas > 0 THEN curtidas - 1 
			ELSE 0 
		END 
		where id = ?`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}
