package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/quangvu30/order-go/config"
)

func Connect() *mongo.Database {
	cfg, _ := config.GetConfig()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURL))
	if err != nil {
		panic(err)
	}
	database := client.Database(cfg.Database)
	return database
}
