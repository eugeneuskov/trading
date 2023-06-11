package repositories

import "trading/app/structures"

type Exchange interface {
	Auth(credentials *structures.Auth) *structures.Token
	SetupOrder(order *structures.Order) error
	CancelOrder(order *structures.Order) error
}
