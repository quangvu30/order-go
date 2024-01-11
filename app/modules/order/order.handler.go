package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/quangvu30/order-go/app/models"
	"github.com/quangvu30/order-go/types"
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
			ctx.JSON(http.StatusUnprocessableEntity, types.Response{
				Code:    400,
				Payload: err,
			})
			return
		}

		orderId := primitive.NewObjectID()
		orderData := models.Order{
			ID:      orderId,
			OrderID: orderId,
		}
		copier.Copy(&orderData, &newOrder)
		order, err := h.Service.Create(orderData)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, types.Response{
				Code:    400,
				Payload: err,
			})
			return
		}
		ctx.JSON(http.StatusCreated, types.Response{
			Code:    200,
			Payload: order,
		})
	}
}

func (h *OrderHandler) GetOrderById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId := ctx.Param("orderId")
		order, err := h.Service.GetOrderById(orderId)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, types.Response{
				Code:    400,
				Payload: err,
			})
			return
		}
		ctx.JSON(http.StatusOK, types.Response{
			Code:    200,
			Payload: order,
		})
	}
}

func (h *OrderHandler) ListOrderByAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accountCode := ctx.Query("accountCode")
		orders, err := h.Service.ListOrderByAccount(accountCode)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, types.Response{
				Code:    400,
				Payload: err,
			})
			return
		}
		ctx.JSON(http.StatusOK, types.Response{
			Code:    200,
			Payload: orders,
		})
	}
}
