package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"trading/app"
	"trading/config"
)

func main() {
	appConfig, err := new(config.Config).Init()
	if err != nil {
		log.Fatalf("Failed initializing config: %s\n", err.Error())
		return
	}

	println("App started")

	application := app.NewApplication(appConfig)
	application.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	println("\nApp shutting down...")

	application.Shutdown()
}
