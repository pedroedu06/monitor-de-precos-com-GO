package model

import (
	"errors"
	"strings"
	"time"
)

type UserDomain struct {
    ID        string    `json:"id"`
    Telefone  string    `json:"telefone"`
    CreatedAt time.Time `json:"created_at"`
}

func (u *UserDomain) Validate() error {
	if u.Telefone == "" {
		return errors.New("telefone obrigatório")
	}
	if !strings.HasPrefix(u.Telefone, "+"){
		return errors.New("Telefone deve ter o codigo do pais!")
	}
	return nil
}