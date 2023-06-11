package structures

import "github.com/google/uuid"

type Order struct {
	id        uuid.UUID
	symbol    string
	side      string
	orderType string
	price     float64
	amount    float64
}

func NewOrder(
	symbol, side, OrderType string,
	price, amount float64,
) *Order {
	return &Order{
		id:        uuid.New(),
		symbol:    symbol,
		side:      side,
		orderType: OrderType,
		price:     price,
		amount:    amount,
	}
}

func (o *Order) Id() uuid.UUID {
	return o.id
}

func (o *Order) Symbol() string {
	return o.symbol
}

func (o *Order) Side() string {
	return o.side
}

func (o *Order) OrderType() string {
	return o.orderType
}

func (o *Order) Price() float64 {
	return o.price
}

func (o *Order) Amount() float64 {
	return o.amount
}
