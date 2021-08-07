package models

import (
	"devbook/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

func (usr *Usuario) Preparar(etapa string) error {
	if erro := usr.validar(etapa); erro != nil {
		return erro
	}

	if erro := usr.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usr *Usuario) validar(etapa string) error {
	if usr.Nome == "" {
		return errors.New("nome obrigatório e não pode estar em branco")
	}

	if usr.Nick == "" {
		return errors.New("nick obrigatório e não pode estar em branco")
	}

	if usr.Email == "" {
		return errors.New("e-mail obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usr.Email); erro != nil {
		return errors.New("o e-mail informado é invalido")
	}

	if etapa == "cadastro" && usr.Senha == "" {
		return errors.New("senha obrigatório e não pode estar em branco")
	}

	return nil
}

func (usr *Usuario) formatar(etapa string) error {
	usr.Nome = strings.TrimSpace(usr.Nome)
	usr.Nick = strings.TrimSpace(usr.Nick)
	usr.Email = strings.TrimSpace(usr.Email)

	if etapa == "cadastro" {
		senhaHash, erro := seguranca.Hash(usr.Senha)
		if erro != nil {
			return erro
		}
		usr.Senha = string(senhaHash)
	}

	return nil
}
