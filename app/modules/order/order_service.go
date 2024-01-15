package order

import (
	"context"
	"fmt"
	"time"

	"github.com/quangvu30/order-go/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	pb "github.com/quangvu30/order-go/grpc"
)

type OrderService struct {
	Ctx        context.Context
	Collection *mongo.Collection
	ClientRPC  pb.OrderTransferClient
}

func NewOrderService(d *mongo.Database, conn *grpc.ClientConn) *OrderService {
	ctx := context.TODO() // context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Collection("orders")
	return &OrderService{
		Collection: collection,
		Ctx:        ctx,
		ClientRPC:  pb.NewOrderTransferClient(conn),
	}
}

func (s *OrderService) Create(order models.Order, modifier string) (orderId string, err error) {
	result, err := s.Collection.InsertOne(s.Ctx, order)
	if err != nil {
		return primitive.NilObjectID.Hex(), err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID.Hex(), fmt.Errorf("InsertedID not correctly converted")
	} else {
		// Transfer the order to pre-order
		response, err := s.ClientRPC.CreateOrder(s.Ctx, &pb.Order{
			OrderID:          order.OrderID.Hex(),
			KindOfOrder:      order.KindOfOrder,
			KindOfOffer:      order.KindOfOffer,
			OrderType:        order.OrderType,
			Commodity:        order.Commodity.Hex(),
			AccountCode:      order.AccountCode,
			Volume1Lot:       order.Volume1Lot,
			DepositRate:      order.DepositRate,
			TransactionType:  order.TransactionType,
			DeliveryTime:     order.DeliveryTime,
			DeliveryLocation: order.DeliveryLocation,
			Price:            order.Price,
			OrderVolume:      order.OrderVolume,
			Packing:          order.Packing,
			Assessor:         order.Assessor,
			OrderValidity:    order.OrderValidity,
			Status:           order.Status,
			CreatedAt:        time.Now().UnixMilli(),
			CreatedBy:        modifier,
		})
		if err != nil {
			return oid.Hex(), err
		}
		return response.OrderID, nil
	}
}

func (s *OrderService) GetOrderById(orderId string) (order models.Order, err error) {
	// Convert id string to ObjectId
	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return models.Order{}, err
	}
	s.Collection.FindOne(s.Ctx, bson.D{{"_id", id}}).Decode(&order)
	return order, nil
}

func (s *OrderService) ListOrderByAccount(accountCode string) (orders []models.Order, err error) {
	cursor, err := s.Collection.Find(s.Ctx, bson.D{{"accountCode", accountCode}})
	if err != nil {
		return nil, err
	}

	// Map result
	if err = cursor.All(s.Ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
