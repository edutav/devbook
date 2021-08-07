package models

import (
	"errors"
	"strings"
	"time"
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

	usr.formatar()
	return nil
}

func (usr *Usuario) validar(etapa string) error {
	if usr.Nome == "" {
		return errors.New("Nome obrigatório e não pode estar em branco")
	}

	if usr.Nick == "" {
		return errors.New("Nick obrigatório e não pode estar em branco")
	}

	if usr.Email == "" {
		return errors.New("Email obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && usr.Senha == "" {
		return errors.New("Senha obrigatório e não pode estar em branco")
	}

	return nil
}

func (usr *Usuario) formatar() {
	usr.Nome = strings.TrimSpace(usr.Nome)
	usr.Nick = strings.TrimSpace(usr.Nick)
	usr.Email = strings.TrimSpace(usr.Email)
}
