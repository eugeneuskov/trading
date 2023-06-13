package app

import (
	"fmt"
	"log"
	"trading/app/repositories/exchage"
	"trading/app/services"
	"trading/config"
)

type Application struct {
	config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{config}
}

func (app *Application) Run() {
	exchanges := exchage.NewExchangeService(
		app.config.Exchanges,
		services.NewHttpClientService(),
	).Exchanges()

	for _, exchange := range exchanges {
		// just for testing
		balances, err := exchange.GetBalances()
		if err != nil {
			log.Fatalln(err.Error())
		}
		for _, balance := range balances.Balances {
			println("coin:", balance.Coin)
			println("free:", fmt.Sprintf("%.8f", balance.Free))
			println("locked:", fmt.Sprintf("%.8f", balance.Locked))
			println("---")
		}
	}
}

func (app *Application) Shutdown() {
	println("Off")
}
