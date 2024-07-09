package app

import (
	"fmt"
	"log"
	"trading/app/repositories/exchage"
	"trading/app/services"
	"trading/app/structures"
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
		println()
		println("------------------------")
		println(exchange.Id())
		println("------------------------")

		// just for testing

			deletedOrder, err := exchange.CancelOrder(
				structures.NewDeleteOrder(
					"28311519742",
					"BTCUSDT",
				),
			)
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Printf("deletedOrder: %+v", deletedOrder)

		newOrder, err := exchange.SetupOrder(
			structures.NewCreateOrder(
				"NOTUSDT",
				"BUY",
				"LIMIT",
				"0.0011",
				"4750",
			),
		)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Printf("newOrder: %+v", newOrder)
		/*
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
		*/
	}
}

func (app *Application) Shutdown() {
	println("Off")
}
