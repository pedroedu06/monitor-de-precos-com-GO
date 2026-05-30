package repository

import (
	"database/sql"

	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserReposity() *UserRepository {
	db := connectoDB()
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u model.UserDomain) (model.UserDomain, error) {
	query := `
		INSERT INTO usuarios (telefone)
        VALUES ($1)
        RETURNING id, telefone, created_at
	`

	var s model.UserDomain
	err := r.db.QueryRow(query, u.Telefone).Scan(
		&s.ID,
		&s.Telefone,
		&s.CreatedAt,
	)
	return s, err
}

func (r *UserRepository) FindByTelefone(telefone string) (model.UserDomain, error) {
    query := `SELECT id, telefone, created_at FROM usuarios WHERE telefone = $1`
    
    var u model.UserDomain
    err := r.db.QueryRow(query, telefone).Scan(&u.ID, &u.Telefone, &u.CreatedAt)
    return u, err
}