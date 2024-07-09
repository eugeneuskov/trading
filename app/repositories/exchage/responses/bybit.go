package responses

// Balance

type ByBitBalanceCoin struct {
	Coin    string `json:"coin"`
	Balance string `json:"walletBalance"`
	Free    string `json:"free"`
	Locked  string `json:"locked"`
}

type ByBitBalanceList struct {
	AccountType string             `json:"accountType"`
	Coins       []ByBitBalanceCoin `json:"coin"`
}

type ByBitAccountBalance struct {
	List []ByBitBalanceList `json:"list"`
}

type ByBitAccountBalanceResponse struct {
	Result ByBitAccountBalance `json:"result"`
}

// Create Order Request

type ByBitCreateOrderRequest struct {
	Category    string `json:"category"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	OrderType   string `json:"orderType"`
	Price       string `json:"price"`
	Quantity    string `json:"qty"`
	TimeInForce string `json:"timeInForce"`
}
