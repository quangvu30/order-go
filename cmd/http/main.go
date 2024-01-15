package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/quangvu30/order-go/app/routes"
	"github.com/quangvu30/order-go/config"
	"github.com/quangvu30/order-go/logger"
	"github.com/quangvu30/order-go/mongo"
)

func main() {
	logging, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
		return
	}
	config, err := config.GetConfig()
	if err != nil {
		logging.Error("Failed to get config", err)
		return
	}
	mongoDatabase := mongo.Connect()

	// Set up a connection to the server.
	conn, err := grpc.Dial(config.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logging.Error("did not connect: ", err)
		return
	}
	defer conn.Close()

	gin.SetMode(config.GinMode)
	engine, err := routes.InitRouter(mongoDatabase, conn, config)
	if err != nil {
		logging.Error("Failed to init gin", err)
		return
	}

	logging.Info("Starting server at http://localhost:", config.ListenPort)
	e := engine.Run(":" + config.ListenPort)
	if e != nil {
		logging.Error("Failed to run server", e)
	}
}
