package structures

type CreateOrder struct {
	symbol    string
	side      string
	orderType string
	price     string
	quantity  string
}

func NewCreateOrder(
	symbol, side, OrderType,
	price, quantity string,
) *CreateOrder {
	return &CreateOrder{
		symbol:    symbol,
		side:      side,
		orderType: OrderType,
		price:     price,
		quantity:  quantity,
	}
}

func (co *CreateOrder) Symbol() string {
	return co.symbol
}

func (co *CreateOrder) Side() string {
	return co.side
}

func (co *CreateOrder) OrderType() string {
	return co.orderType
}

func (co *CreateOrder) Price() string {
	return co.price
}

func (co *CreateOrder) Quantity() string {
	return co.quantity
}

type DeleteOrder struct {
	id     string
	symbol string
}

func NewDeleteOrder(id, symbol string) *DeleteOrder {
	return &DeleteOrder{
		id:     id,
		symbol: symbol,
	}
}

func (do *DeleteOrder) ID() string  {
	return do.id
}

func (do *DeleteOrder) Symbol() string {
	return do.symbol
}
