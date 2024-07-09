package repositories

import "trading/app/structures"

type Exchange interface {
	Id() string
	Auth() *structures.Token
	GetBalances() (*structures.BalanceInfo, error)
	SetupOrder(order *structures.CreateOrder) (*structures.ExchangeOrder, error)
	CancelOrder(order *structures.DeleteOrder) (*structures.ExchangeOrder, error)
}
