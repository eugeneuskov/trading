package exchage

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"trading/app/repositories/exchage/responses"
	"trading/app/services"
	"trading/app/structures"
)

const (
	binanceSignatureKey = "signature"

	binanceUriAccount      = "/api/v3/account"
	binanceUriCreateOrder  = "/api/v3/order"
	binanceUriDeleteOrders = "/api/v3/order"
)

type BinanceExchange struct {
	id          string
	credentials structures.Auth
	baseUrls    []string
	httpClient  *services.HttpClientService
}

func NewBinanceExchange(
	id string,
	credentials structures.Auth,
	baseUrls []string,
	httpClient *services.HttpClientService,
) *BinanceExchange {
	return &BinanceExchange{
		id:          id,
		credentials: credentials,
		baseUrls:    baseUrls,
		httpClient:  httpClient,
	}
}

func (e *BinanceExchange) Id() string {
	return e.id
}

func (e *BinanceExchange) Auth() *structures.Token {
	return &structures.Token{}
}

func (e *BinanceExchange) GetBalances() (*structures.BalanceInfo, error) {
	req, err := e.buildRequest(binanceUriAccount, url.Values{})
	if err != nil {
		return nil, err
	}

	responseBody, err := e.httpClient.Get(req.fullUrl, req.header, req.query)
	if err != nil {
		return nil, err
	}

	var balanceData responses.BinanceBalanceInfoResponse
	err = json.Unmarshal(responseBody, &balanceData)
	if err != nil {
		return nil, err
	}

	balances := make([]structures.Balance, 0)
	for _, balance := range balanceData.Balances {
		balanceFree, _ := strconv.ParseFloat(balance.Free, 64)
		balanceLocked, _ := strconv.ParseFloat(balance.Locked, 64)
		if balanceFree > 0.0 || balanceLocked > 0.0 {
			balances = append(balances, structures.Balance{
				Coin:   balance.Coin,
				Free:   balanceFree,
				Locked: balanceLocked,
			})
		}
	}

	return &structures.BalanceInfo{
		CanTrade:    balanceData.CanTrade,
		CanWithdraw: balanceData.CanWithdraw,
		CanDeposit:  balanceData.CanDeposit,
		Balances:    balances,
	}, nil
}

func (e *BinanceExchange) SetupOrder(order *structures.CreateOrder) (*structures.ExchangeOrder, error) {
	queryParams := url.Values{}
	queryParams.Add("symbol", order.Symbol())
	queryParams.Add("side", order.Side())
	queryParams.Add("type", order.OrderType())
	queryParams.Add("timeInForce", "GTC")
	queryParams.Add("quantity", order.Quantity())
	queryParams.Add("price", order.Price())

	req, err := e.buildRequest(binanceUriCreateOrder, queryParams)
	if err != nil {
		return nil, err
	}

	responseBody, err := e.httpClient.Post(req.fullUrl, req.header, []byte(req.query.Encode()))
	if err != nil {
		return nil, err
	}

	var newOrder responses.BinanceOrderResponse
	err = json.Unmarshal(responseBody, &newOrder)
	if err != nil {
		return nil, err
	}

	return &structures.ExchangeOrder{
		Symbol:       newOrder.Symbol,
		OrderId:      newOrder.OrderId,
		TransactTime: newOrder.TransactTime,
		Price:        newOrder.Price,
		Quantity:     newOrder.OriginalQuantity,
		Status:       newOrder.Status,
		Type:         newOrder.Type,
		Side:         newOrder.Side,
	}, nil
}

func (e *BinanceExchange) CancelOrder(order *structures.DeleteOrder) (*structures.ExchangeOrder, error) {
	queryParams := url.Values{}
	queryParams.Set("symbol", order.Symbol())
	queryParams.Set("orderId", order.ID())

	req, err := e.buildRequest(binanceUriDeleteOrders, queryParams)
	if err != nil {
		return nil, err
	}

	responseBody, err := e.httpClient.Delete(req.fullUrl, req.header, req.query)
	if err != nil {
		return nil, err
	}

	var deletedOrder responses.BinanceOrderResponse
	err = json.Unmarshal(responseBody, &deletedOrder)
	if err != nil {
		return nil, err
	}

	return &structures.ExchangeOrder{
			Symbol:       deletedOrder.Symbol,
			OrderId:      deletedOrder.OrderId,
			TransactTime: deletedOrder.TransactTime,
			Price:        deletedOrder.Price,
			Quantity:     deletedOrder.OriginalQuantity,
			Status:       deletedOrder.Status,
			Type:         deletedOrder.Type,
			Side:         deletedOrder.Side,
		}, nil
}

func (e *BinanceExchange) buildRequest(
	endpoint string,
	query url.Values,
) (*request, error) {
	query.Set(timestampKey, strconv.FormatInt(time.Now().Unix()*1000, 10))
	queryString := query.Encode()
	query.Set(binanceSignatureKey, e.signRequest(queryString))

	header := http.Header{}
	header.Set("X-MBX-APIKEY", e.credentials.ApiKey())
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	return &request{
		query:   query,
		header:  header,
		fullUrl: fmt.Sprintf("%s%s", e.baseUrls[0], endpoint),
	}, nil
}

func (e *BinanceExchange) signRequest(raw string) string {
	h := hmac.New(sha256.New, []byte(e.credentials.ApiSecret()))
	h.Write([]byte(raw))

	return hex.EncodeToString(h.Sum(nil))
}
