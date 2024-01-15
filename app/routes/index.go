package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quangvu30/order-go/app/middlewares"
	"github.com/quangvu30/order-go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func InitRouter(db *mongo.Database, connGrpc *grpc.ClientConn, cfg *config.Config) (*gin.Engine, error) {
	engine := gin.Default()
	middle := middlewares.NewAuthMiddleware(cfg)
	v1 := engine.Group("/api/v1")
	{
		addOrderRouter(db, v1, middle, connGrpc)
	}

	return engine, nil
}
