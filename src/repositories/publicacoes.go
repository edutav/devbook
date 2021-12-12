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
