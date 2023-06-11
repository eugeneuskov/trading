package structures

type Wallet struct {
	address string
	coin    string
	balance float64
}

func NewWallet(address string, coin string) *Wallet {
	return &Wallet{
		address: address,
		coin:    coin,
	}
}

func (w *Wallet) Address() string {
	return w.address
}

func (w *Wallet) Coin() string {
	return w.coin
}

func (w *Wallet) setBalance(balance float64) {
	w.balance = balance
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
