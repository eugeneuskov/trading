package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"trading/config"
)

func main() {
	appConfig, err := new(config.Config).Init()
	if err != nil {
		log.Fatalf("Failed initializing config: %s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", appConfig)

	println("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	println("\nApp shutting down...")
}
