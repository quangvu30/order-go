package order

import (
	"context"
	"fmt"

	"github.com/quangvu30/order-go/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	Ctx        context.Context
	Collection *mongo.Collection
}

func NewOrderService(d *mongo.Database) *OrderService {
	ctx := context.TODO() // context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Collection("orders")
	return &OrderService{
		Collection: collection,
		Ctx:        ctx,
	}
}

func (s *OrderService) Create(order models.Order) (orderId primitive.ObjectID, err error) {
	result, err := s.Collection.InsertOne(s.Ctx, order)
	if err != nil {
		return primitive.NilObjectID, err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("InsertedID not correctly converted")
	} else {
		return oid, nil
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
