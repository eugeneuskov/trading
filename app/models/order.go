package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id              uuid.UUID
	Symbol          string
	ExchangerId     string    `db:"exchanger_id"`
	ExchangeOrderId string    `db:"exchange_order_id"`
	CreatedAt       time.Time `db:"created_at"`
}
