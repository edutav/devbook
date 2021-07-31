package database

import (
	"database/sql"
	"devbook/src/config"

	_ "github.com/go-sql-driver/mysql"
)

// Abre a conexão com o banco de dados e retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysq", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
