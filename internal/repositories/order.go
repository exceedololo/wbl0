package repositories

import (
	"bwTechLvl0/internal/database"
	"bwTechLvl0/internal/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"

	//"database/sql"
	"encoding/json"
	"time"
)

type OrderRepo struct {
	db *database.DataBase
}

func NewOrderRepo( /*ctx context.Context,*/ db *database.DataBase) (*OrderRepo, error) {
	/*newCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	for {
		select {
		case <-newCtx.Done():
			return nil, ctx.Err()
		default:
			return &OrderRepo{db: db}, nil
		}
	}*/
	//or i could use this
	//ctx, cancel:= context.WithTimeout(context.Background(), 15 * time.Second)
	//defer cancel()
	return &OrderRepo{db: db}, nil

}

//upsert, getById

// on conflict, сделать чтоб норм работало с базой, без коллизий
func (or *OrderRepo) Upsert(ctx context.Context, order models.Order) error {
	orderData, err := json.Marshal(order.Data)
	if err != nil {
		return err
	}

	// Создаем контекст с установленным сроком действия
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// Используем контекст для выполнения запроса
	_, err = or.db.Conn.Exec(ctxWithTimeout,
		"INSERT INTO orders (order_uid, date_created, data)"+
			"VALUES($1, $2, $3)",
		order.OrderUID, order.DateCreated, orderData,
	)
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepo) GetById(ctx context.Context, orderUID string) (*models.Order, error) {
	/*select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		row := or.db.Conn.QueryRowContext(ctx,
			"SELECT order_uid, date_created, data FROM orders WHERE order_uid = $1",
			orderUID,
		)*/
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
		return nil, err
	}
	order.Data = orderData
	return &order, nil
}

//}
