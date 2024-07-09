package exchage

import (
	"net/http"
	"net/url"
	"trading/app/repositories"
	"trading/app/services"
	"trading/app/structures"
	"trading/config"
)

const (
	exchangeBinanceId = "binance"
	exchangeBybitId   = "bybit"

	timestampKey = "timestamp"
)

type ExchangeService struct {
	exchangeConfig []config.Exchange
	httpClient     *services.HttpClientService
}

func NewExchangeService(
	exchangeConfig []config.Exchange,
	httpClient *services.HttpClientService,
) *ExchangeService {
	return &ExchangeService{
		exchangeConfig,
		httpClient,
	}
}

func (s *ExchangeService) Exchanges() []repositories.Exchange {
	exchanges := make([]repositories.Exchange, 0, len(s.exchangeConfig))

	for _, conf := range s.exchangeConfig {
		switch conf.Id {
		case exchangeBinanceId:
			exchanges = append(
				exchanges,
				NewBinanceExchange(
					conf.Id,
					structures.NewAuth(conf.ApiKey, conf.ApiSecret),
					conf.Url,
					s.httpClient,
				),
			)
		case exchangeBybitId:
			exchanges = append(
				exchanges,
				NewByBitExchange(
					conf.Id,
					structures.NewAuth(conf.ApiKey, conf.ApiSecret),
					conf.Url,
					s.httpClient,
				),
			)
		}
	}

	return exchanges
}

type request struct {
	fullUrl string
	query   url.Values
	header  http.Header
}
