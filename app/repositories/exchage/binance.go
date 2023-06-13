package exchage

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"trading/app/services"
	"trading/app/structures"
)

const (
	timestampKey = "timestamp"
	signatureKey = "signature"
)

type balanceResponse struct {
	Coin   string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type balanceInfoResponse struct {
	CanTrade    bool              `json:"canTrade"`
	CanWithdraw bool              `json:"canWithdraw"`
	CanDeposit  bool              `json:"canDeposit"`
	Balances    []balanceResponse `json:"balances"`
}

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
	req, err := e.request("/api/v3/account", url.Values{})
	if err != nil {
		return nil, err
	}

	responseBody, err := e.httpClient.Get(req.fullUrl, req.header, req.query)
	if err != nil {
		return nil, err
	}

	var balanceData balanceInfoResponse
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

func (e *BinanceExchange) SetupOrder(order *structures.Order) error {
	return nil
}

func (e *BinanceExchange) CancelOrder(order *structures.Order) error {
	return nil
}

type request struct {
	method     string
	endpoint   string
	query      url.Values
	form       url.Values
	recvWindow int64
	secType    int8
	header     http.Header
	body       io.Reader
	fullUrl    string
}

func (e *BinanceExchange) request(
	endpoint string,
	query url.Values,
) (*request, error) {
	query.Set(timestampKey, fmt.Sprintf("%v", time.Now().UnixNano()/int64(time.Millisecond)))
	queryString := query.Encode()

	signature, err := e.signRequest(queryString)
	if err != nil {
		return nil, err
	}
	query.Set(signatureKey, signature)

	header := http.Header{}
	header.Set("X-MBX-APIKEY", e.credentials.ApiKey())

	return &request{
		endpoint:   endpoint,
		query:      query,
		form:       nil,
		recvWindow: 0,
		secType:    2,
		header:     header,
		body:       nil,
		fullUrl:    fmt.Sprintf("%s%s", e.baseUrls[0], endpoint),
	}, nil
}

func (e *BinanceExchange) signRequest(raw string) (string, error) {
	mac := hmac.New(sha256.New, []byte(e.credentials.ApiSecret()))
	_, err := mac.Write([]byte(raw))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}
