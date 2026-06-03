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

func (r *ProdutoRepository) ListByUserID(userID string) ([]model.ProdutoDomain, error) {
	query := `
		SELECT id, usuario_id, url, preco_alvo, ativo
		FROM produtos
		WHERE usuario_id = $1
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []model.ProdutoDomain
	for rows.Next() {
		var p model.ProdutoDomain
		if err := rows.Scan(&p.ID, &p.UserID, &p.URL, &p.TargetPrice, &p.IsActive); err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return produtos, nil
}

// ListAllActive returns every active product across all users.
// Used by the background worker to know which products to scrape.
func (r *ProdutoRepository) ListAllActive() ([]model.ProdutoDomain, error) {
	query := `
		SELECT id, usuario_id, url, preco_alvo, ativo
		FROM produtos
		WHERE ativo = true
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.ProdutoDomain
	for rows.Next() {
		var p model.ProdutoDomain
		if err := rows.Scan(&p.ID, &p.UserID, &p.URL, &p.TargetPrice, &p.IsActive); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProdutoRepository) DeleteProductById(id, userId string) (int64, error) {
	query := `DELETE FROM produtos 
		WHERE id = $1 AND usuario_id = $2`

	result, err := r.db.Exec(query, id, userId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

	return rowsAffected, nil
}

func (r *ProdutoRepository) SetActiveById(id, userID string, active bool) (model.ProdutoDomain, error) {
	query := `
		UPDATE produtos
		SET ativo = $3
		WHERE id = $1 AND usuario_id = $2
		RETURNING id, usuario_id, url, preco_alvo, ativo
	`
	var p model.ProdutoDomain
	err := r.db.QueryRow(query, id, userID, active).Scan(
		&p.ID, &p.UserID, &p.URL, &p.TargetPrice, &p.IsActive,
	)
	return p, err
}
