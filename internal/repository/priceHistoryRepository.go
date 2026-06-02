package repository

import (
	"database/sql"

	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
)

type PriceHistoryRepository struct {
	db *sql.DB
}

func NewPriceHistoryRepository() *PriceHistoryRepository {
	db := connectoDB()
	return &PriceHistoryRepository{db: db}
}

// InsertPrice stores a new price entry for the given product.
// Used by the worker after capturing a new (changed) price.
func (r *PriceHistoryRepository) InsertPrice(productID string, price float64) (model.PriceHistoryDomain, error) {
	query := `
		INSERT INTO historico_precos (produto_id, preco)
		VALUES ($1, $2)
		RETURNING id, produto_id, preco, capturado_em
	`
	var h model.PriceHistoryDomain
	err := r.db.QueryRow(query, productID, price).Scan(
		&h.ID, &h.ProductID, &h.Price, &h.CapturedAt,
	)
	return h, err
}

// GetLatestPrice returns the most recent price entry for the product.
// Used by the worker to decide whether the captured price is new.
// Returns sql.ErrNoRows if the product has no price history yet.
func (r *PriceHistoryRepository) GetLatestPrice(productID string) (model.PriceHistoryDomain, error) {
	query := `
		SELECT id, produto_id, preco, capturado_em
		FROM historico_precos
		WHERE produto_id = $1
		ORDER BY capturado_em DESC
		LIMIT 1
	`
	var h model.PriceHistoryDomain
	err := r.db.QueryRow(query, productID).Scan(
		&h.ID, &h.ProductID, &h.Price, &h.CapturedAt,
	)
	return h, err
}

// ListByProductID returns the full price history for a product,
// but only if the product belongs to the given user.
// The JOIN with produtos enforces ownership — if the product is not the user's,
// the result is an empty slice (no leakage about whether the product exists).
func (r *PriceHistoryRepository) ListByProductID(productID, userID string) ([]model.PriceHistoryDomain, error) {
	query := `
		SELECT h.id, h.produto_id, h.preco, h.capturado_em
		FROM historico_precos h
		INNER JOIN produtos p ON p.id = h.produto_id
		WHERE h.produto_id = $1 AND p.usuario_id = $2
		ORDER BY h.capturado_em ASC
	`
	rows, err := r.db.Query(query, productID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []model.PriceHistoryDomain
	for rows.Next() {
		var h model.PriceHistoryDomain
		if err := rows.Scan(&h.ID, &h.ProductID, &h.Price, &h.CapturedAt); err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return history, nil
}
