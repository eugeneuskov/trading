package exchage

import (
	"trading/app/repositories"
	"trading/app/services"
	"trading/app/structures"
	"trading/config"
)

const (
	exchangeBinanceId = "binance"
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
		}
	}

	return exchanges
}
