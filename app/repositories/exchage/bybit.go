package exchage

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
	"trading/app/repositories/exchage/responses"
	"trading/app/services"
	"trading/app/structures"
)

const (
	bybitApiKey       = "api_key"
	bybitSignatureKey = "sign"

	bybitUriAccountBalance = "/v5/account/wallet-balance"
	bybitUriCreateOrder    = "/v5/order/create"
)

type ByBitExchange struct {
	id          string
	credentials structures.Auth
	baseUrls    []string
	httpClient  *services.HttpClientService
}

func NewByBitExchange(
	id string,
	credentials structures.Auth,
	baseUrls []string,
	httpClient *services.HttpClientService,
) *ByBitExchange {
	return &ByBitExchange{
		id:          id,
		credentials: credentials,
		baseUrls:    baseUrls,
		httpClient:  httpClient,
	}
}

func (b *ByBitExchange) Id() string {
	return b.id
}

func (b *ByBitExchange) Auth() *structures.Token {
	return &structures.Token{}
}

func (b *ByBitExchange) GetBalances() (*structures.BalanceInfo, error) {
	queryParams := url.Values{}
	queryParams.Set(bybitApiKey, b.credentials.ApiKey())
	queryParams.Set(timestampKey, strconv.FormatInt(time.Now().Unix()*1000, 10))
	queryParams.Add("accountType", "SPOT")

	req, err := b.buildRequest(bybitUriAccountBalance, queryParams)
	if err != nil {
		return nil, err
	}
	req.query.Set(bybitSignatureKey, b.signRequest(queryParams.Encode()))

	responseBody, err := b.httpClient.Get(req.fullUrl, req.header, req.query)
	if err != nil {
		return nil, err
	}

	var balanceData responses.ByBitAccountBalanceResponse
	err = json.Unmarshal(responseBody, &balanceData)
	if err != nil {
		return nil, err
	}

	balances := make([]structures.Balance, 0)
	for _, list := range balanceData.Result.List {
		for _, balance := range list.Coins {
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
	}

	return &structures.BalanceInfo{
		Balances: balances,
	}, nil
}

func (b *ByBitExchange) SetupOrder(order *structures.CreateOrder) (*structures.ExchangeOrder, error) {
	req, err := b.buildRequest(bybitUriCreateOrder, url.Values{})
	if err != nil {
		return nil, err
	}

	orderParams := map[string]string{
		"category":      "linear",
		"api_key":       b.credentials.ApiKey(),
		"symbol":        order.Symbol(),
		"side":          order.Side(),
		"order_type":    order.OrderType(),
		"qty":           order.Quantity(),
		"price":         order.Price(),
		"time_in_force": "GTC",
		"timestamp":     strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	queryString := createQueryString(orderParams)
	signature := b.signRequest(queryString)
	orderParams[bybitSignatureKey] = signature

	jsonData, err := json.Marshal(orderParams)
	if err != nil {
		return nil, err
	}

	responseBody, err := b.httpClient.Post(
		req.fullUrl,
		req.header,
		jsonData,
	)

	if err != nil {
		return nil, err
	}

	println(string(responseBody))
	return &structures.ExchangeOrder{}, nil
}

func (b *ByBitExchange) CancelOrder(order *structures.DeleteOrder) (*structures.ExchangeOrder, error) {
	return nil, nil
}

func (b *ByBitExchange) buildRequest(
	endpoint string,
	query url.Values,
) (*request, error) {
	header := http.Header{}
	header.Set("Content-Type", "application/json")

	return &request{
		query:   query,
		header:  header,
		fullUrl: fmt.Sprintf("%s%s", b.baseUrls[0], endpoint),
	}, nil
}

func (b *ByBitExchange) signRequest(raw string) string {
	h := hmac.New(sha256.New, []byte(b.credentials.ApiSecret()))
	h.Write([]byte(raw))

	return hex.EncodeToString(h.Sum(nil))
}

func createQueryString(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	queryString := ""
	for _, key := range keys {
		if queryString != "" {
			queryString += "&"
		}
		queryString += key + "=" + params[key]
	}
	return queryString
}
