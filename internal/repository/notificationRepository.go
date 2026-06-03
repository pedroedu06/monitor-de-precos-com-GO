package repository

import (
	"database/sql"

	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository() *NotificationRepository {
	db := connectoDB()
	return &NotificationRepository{db: db}
}

// Create records a notification that was sent to the user's WhatsApp.
// Called by the worker right after the WhatsApp provider confirms delivery.
func (r *NotificationRepository) Create(n model.NotificationDomain) (model.NotificationDomain, error) {
	query := `
		INSERT INTO notificacoes (produto_id, preco_anterior, preco_atual)
		VALUES ($1, $2, $3)
		RETURNING id, produto_id, preco_anterior, preco_atual, enviada_em
	`
	var out model.NotificationDomain
	err := r.db.QueryRow(query, n.ProductID, n.PreviousPrice, n.CurrentPrice).Scan(
		&out.ID,
		&out.ProductID,
		&out.PreviousPrice,
		&out.CurrentPrice,
		&out.SentAt,
	)
	return out, err
}

// ListByUserID returns every notification belonging to the user, newest first.
// The JOIN with produtos enforces ownership: notifications from products
// that don't belong to the user are filtered out at the SQL level.
func (r *NotificationRepository) ListByUserID(userID string) ([]model.NotificationDomain, error) {
	query := `
		SELECT n.id, n.produto_id, n.preco_anterior, n.preco_atual, n.enviada_em
		FROM notificacoes n
		INNER JOIN produtos p ON p.id = n.produto_id
		WHERE p.usuario_id = $1
		ORDER BY n.enviada_em DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []model.NotificationDomain
	for rows.Next() {
		var n model.NotificationDomain
		if err := rows.Scan(
			&n.ID,
			&n.ProductID,
			&n.PreviousPrice,
			&n.CurrentPrice,
			&n.SentAt,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notifications, nil
}
