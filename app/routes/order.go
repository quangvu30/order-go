package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quangvu30/order-go/app/middlewares"
	"github.com/quangvu30/order-go/app/modules/order"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func addOrderRouter(db *mongo.Database, rg *gin.RouterGroup, mid *middlewares.AuthMiddleware, connGrpc *grpc.ClientConn) {
	orderService := order.NewOrderService(db, connGrpc)
	orderHandler := order.NewOrderHandler(orderService)

	orderRouter := rg.Group("/order")
	orderRouter.Use(mid.User())
	orderRouter.POST("/", orderHandler.CreateOrder())
	orderRouter.GET("/:orderId", orderHandler.GetOrderById())
	orderRouter.GET("/", orderHandler.ListOrderByAccount())
}
