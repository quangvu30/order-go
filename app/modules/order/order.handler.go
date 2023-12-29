package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/quangvu30/order-go/app/models"
	"github.com/quangvu30/order-go/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHandler struct {
	Service *OrderService
}

func NewOrderHandler(s *OrderService) *OrderHandler {
	return &OrderHandler{
		Service: s,
	}
}

func (h *OrderHandler) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newOrder ReqCreateOrder
		if err := ctx.BindJSON(&newOrder); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusUnprocessableEntity, config.Response{
				Code:    400,
				Payload: err,
			})
			return
		}

		orderData := models.Order{
			OrderID: primitive.NewObjectID(),
		}
		copier.Copy(&orderData, &newOrder)
		order, err := h.Service.Create(orderData)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, config.Response{
				Code:    400,
				Payload: err,
			})
			return
		}
		ctx.JSON(http.StatusCreated, config.Response{
			Code:    200,
			Payload: order,
		})
	}
}
