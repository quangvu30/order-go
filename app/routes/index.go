package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(db *mongo.Database) (*gin.Engine, error) {
	engine := gin.Default()
	v1 := engine.Group("/api/v1")
	{
		addOrderRouter(db, v1)
	}

	return engine, nil
}
