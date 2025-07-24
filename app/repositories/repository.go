package repositories

import (
	"github.com/jmoiron/sqlx"
	"trading/app/models"
)

const (
	orderTableName  = "orders"
	signalTableName = "signals"

	symbolColumn    = "symbol"
	createdAtColumn = "created_at"
)

type Order interface {
	Insert(symbol, exchangerId, exchangerOrderId string) (*models.Order, error)
	GetBySymbol(symbol, exchangerId string) ([]*models.Order, error)
	Delete(symbol, exchangerId string) error
}

type Signal interface {
	Insert(symbol, value string) (*models.Signal, error)
	GetBySymbol(symbol string) (*models.Signal, error)
	Update(signal *models.Signal, value string) (*models.Signal, error)
}

type Repository struct {
	Order
	Signal
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order:  newOrderRepository(db),
		Signal: newSignalRepository(db),
	}
}
