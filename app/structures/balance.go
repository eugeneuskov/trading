package structures

type Balance struct {
	Coin   string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

type BalanceInfo struct {
	CanTrade    bool      `json:"canTrade"`
	CanWithdraw bool      `json:"canWithdraw"`
	CanDeposit  bool      `json:"canDeposit"`
	Balances    []Balance `json:"balances"`
}
