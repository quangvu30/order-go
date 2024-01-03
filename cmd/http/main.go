package main

import (
	"log"

	"github.com/quangvu30/order-go/app/routes"
	"github.com/quangvu30/order-go/config"
	"github.com/quangvu30/order-go/mongo"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Panic("Failed to get config", err)
	}

	mongoDatabase := mongo.Connect()
	engine, err := routes.InitRouter(mongoDatabase)
	if err != nil {
		log.Panic("Failed to init gin", err)
	}

	_ = engine.Run(":" + config.ListenPort)
	log.Println("Server is running on port " + config.ListenPort)
}
