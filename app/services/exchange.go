package services

import "trading/app/structures"

type Exchange interface {
	Id() string
	GetBalances() (*structures.BalanceInfo, error)
	SetupOrder(order *structures.CreateOrder) (*structures.ExchangeOrder, error)
	CancelOrder(order *structures.DeleteOrder) (*structures.ExchangeOrder, error)
}
