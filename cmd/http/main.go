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
	_, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
		return
	}
	config, err := config.GetConfig()
	if err != nil {
		logger.Log.Error("Failed to get config", err.Error())
		return
	}
	mongoDatabase := mongo.Connect()

	// Set up a connection to the server.
	conn, err := grpc.Dial(config.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Error("did not connect: ", err.Error())
		return
	}
	defer conn.Close()

	gin.SetMode(config.GinMode)
	engine, err := routes.InitRouter(mongoDatabase, conn, config)
	if err != nil {
		logger.Log.Error("Failed to init gin", err.Error())
		return
	}

	logger.Log.Info("Starting server at http://localhost:", config.ListenPort)
	e := engine.Run(":" + config.ListenPort)
	if e != nil {
		logger.Log.Error("Failed to run server", e.Error())
	}
}
