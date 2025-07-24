package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"trading/app/repositories"
	"trading/app/services"
	"trading/app/services/exchage"
	"trading/config"
)

type Application struct {
	config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{config}
}

func (app *Application) Shutdown() {
	println("Off")
}

func (app *Application) Run() {
	db, err := app.newPostgresDb()
	if err != nil {
		log.Fatalf("Failed to initialize DB: %s\n", err.Error())
		return
	}
	reps := repositories.NewRepository(db)

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

		orders, err := reps.Order.GetBySymbol(
			"USDTBTC",
			exchange.Id(),
		)
		if err != nil {
			log.Fatalf("Failed to create order: %s\n", err.Error())
		}
		for _, order := range orders {
			fmt.Printf("%+v\n", order)
		}

		/*
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

func (app *Application) newPostgresDb() (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			app.config.Db.Host,
			app.config.Db.Port,
			app.config.Db.User,
			app.config.Db.DbName,
			app.config.Db.Password,
			app.config.Db.SslMode,
		),
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
