package repositories

import "trading/app/structures"

type Robot interface {
	StartStrategy() error
	SetupOrder(order *structures.CreateOrder) error
	CancelOrder(order *structures.CreateOrder) error
}
