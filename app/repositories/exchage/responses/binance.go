package responses

// Balance

type BinanceBalanceResponse struct {
	Coin   string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type BinanceBalanceInfoResponse struct {
	CanTrade    bool                     `json:"canTrade"`
	CanWithdraw bool                     `json:"canWithdraw"`
	CanDeposit  bool                     `json:"canDeposit"`
	Balances    []BinanceBalanceResponse `json:"balances"`
}

// Order

type BinanceOrderResponse struct {
	Symbol                   string `json:"symbol"`
	OrderId                  uint   `json:"orderId"`
	OrderListId              int    `json:"orderListId"`
	ClientOrderId            string `json:"clientOrderId"`
	TransactTime             uint   `json:"transactTime"`
	Price                    string `json:"price"`
	OriginalQuantity         string `json:"origQty"`
	ExecutedQuantity         string `json:"executedQty"`
	CummulativeQuoteQuantity string `json:"cummulativeQuoteQty"`
	Status                   string `json:"status"`
	TimeInForce              string `json:"timeInForce"`
	Type                     string `json:"type"`
	Side                     string `json:"side"`
	WorkingTime              uint   `json:"workingTime"`
	SelfTradePreventionMode  string `json:"selfTradePreventionMode"`
}
