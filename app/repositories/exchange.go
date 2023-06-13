package repositories

import "trading/app/structures"

type Exchange interface {
	Id() string
	Auth() *structures.Token
	GetBalances() (*structures.BalanceInfo, error)
	SetupOrder(order *structures.Order) error
	CancelOrder(order *structures.Order) error
}
