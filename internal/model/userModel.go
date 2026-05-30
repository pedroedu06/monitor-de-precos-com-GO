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
		return errors.New("phone is required")
	}
	if !strings.HasPrefix(u.Telefone, "+"){
		return errors.New("phone must include the country code!")
	}
	return nil
}