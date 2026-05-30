package repository

import (
	"database/sql"

	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
)

type ProdutoRepository struct {
	db *sql.DB
}

func NewProdutoRepository() *ProdutoRepository {
	db := connectoDB()
	return &ProdutoRepository{db: db}
}

func (r *ProdutoRepository) Create(produto model.ProdutoDomain) (model.ProdutoDomain, error) {
	query := `
        INSERT INTO produtos (usuario_id, url, preco_alvo)
        VALUES ($1, $2, $3)
        RETURNING id, usuario_id, url, preco_alvo, ativo
    `
	var p model.ProdutoDomain
	err := r.db.QueryRow(query, produto.UserID, produto.URL, produto.TargetPrice).Scan(
		&p.ID, &p.UserID, &p.URL, &p.TargetPrice, &p.IsActive,
	)
	return p, err
}
