package repositories

import (
	"bwTechLvl0/internal/database"
	"bwTechLvl0/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"

	"encoding/json"
)

type OrderRepo struct {
	db *database.DataBase
}

func NewOrderRepo( /*ctx context.Context,*/ db *database.DataBase) (*OrderRepo, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	return &OrderRepo{db: db}, nil

}

//upsert, getById

// on conflict, сделать чтоб норм работало с базой, без коллизий
func (or *OrderRepo) Upsert(ctx context.Context, order models.Order) error {
	orderData, err := json.Marshal(order.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal order data: %w", err)
	}

	var exists bool
	err = or.db.Conn.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM orders WHERE order_uid = $1)", order.OrderUID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if order exists: %w", err)
	}

	if exists {
		// Обновление существующей записи
		_, err := or.db.Conn.Exec(ctx,
			"UPDATE orders SET date_created = $2, data = $3 WHERE order_uid = $1",
			order.OrderUID, order.DateCreated, orderData,
		)
		if err != nil {
			return fmt.Errorf("failed to update order: %w", err)
		}
	} else {
		// Вставка новой записи
		_, err := or.db.Conn.Exec(ctx,
			"INSERT INTO orders (order_uid, date_created, data) VALUES ($1, $2, $3)",
			order.OrderUID, order.DateCreated, orderData,
		)
		if err != nil {
			return fmt.Errorf("failed to insert order: %w", err)
		}
	}

	return nil
}

func (or *OrderRepo) GetById(ctx context.Context, orderUID string) (*models.Order, error) {
	var order models.Order
	var orderData json.RawMessage

	err := or.db.Conn.QueryRow(ctx,
		"SELECT order_uid, date_created, data FROM orders WHERE order_uid= $1",
		orderUID,
	).Scan(&order.OrderUID, &order.DateCreated, &orderData)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil //there is no order with this uid
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	order.Data = orderData
	return &order, nil
}

/*
func (or *OrderRepo) Close() {
	or.db.Conn.Close()
}*/

//}
