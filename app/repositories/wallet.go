package repositories

import "trading/app/structures"

type Wallet interface {
	GetBalance(wallet structures.Wallet)
}
