package repositories

import "trading/app/structures"

type Robot interface {
	StartStrategy() error
	SetupOrder(order *structures.Order) error
	CancelOrder(order *structures.Order) error
}
