package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
	"trading/app/models"
)

const (
	orderIdColumn              = "id"
	orderExchangedIdColumn     = "exchanger_id"
	orderExchangeOrderIdColumn = "exchange_order_id"
)

type orderRepository struct {
	db *sqlx.DB
}

func newOrderRepository(db *sqlx.DB) *orderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) Insert(symbol, exchangerId, exchangerOrderId string) (*models.Order, error) {
	id := uuid.New()
	createdAt := time.Now()

	_, err := o.db.Exec(fmt.Sprintf(`
			insert into %s
			(%s, %s, %s, %s, %s)
			values ($1, $2, $3, $4, $5)
		`,
		orderTableName,
		orderIdColumn,
		symbolColumn,
		orderExchangedIdColumn,
		orderExchangeOrderIdColumn,
		createdAtColumn,
	),
		id,
		symbol,
		exchangerId,
		exchangerOrderId,
		createdAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id:              id,
		Symbol:          symbol,
		ExchangerId:     exchangerId,
		ExchangeOrderId: exchangerOrderId,
		CreatedAt:       createdAt,
	}, nil
}

func (o *orderRepository) GetBySymbol(symbol, exchangerId string) ([]*models.Order, error) {
	orders := make([]*models.Order, 0)

	err := o.db.Select(
		&orders,
		fmt.Sprintf(`
			select * from %s 
			where %s = $1 and %s = $2
		`,
			orderTableName,
			symbolColumn,
			orderExchangedIdColumn,
		),
		symbol,
		exchangerId,
	)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderRepository) Delete(symbol, exchangerId string) error {
	_, err := o.db.Exec(fmt.Sprintf(`
			delete from %s
			where %s = $1 and %s = $2
		`,
		orderTableName,
		symbolColumn,
		orderExchangedIdColumn,
	),
		symbol,
		exchangerId,
	)

	return err
}
