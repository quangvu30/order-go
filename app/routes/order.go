package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quangvu30/order-go/app/modules/order"
	"go.mongodb.org/mongo-driver/mongo"
)

func addOrderRouter(db *mongo.Database, rg *gin.RouterGroup) {
	orderService := order.NewOrderService(db)
	orderHandler := order.NewOrderHandler(orderService)

	orderRouter := rg.Group("/order")
	orderRouter.POST("/", orderHandler.CreateOrder())
	orderRouter.GET("/:orderId", orderHandler.GetOrderById())
	orderRouter.GET("/", orderHandler.ListOrderByAccount())
}
