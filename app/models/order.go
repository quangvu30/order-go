package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	OrderID          primitive.ObjectID `bson:"orderId, omitempty"`
	KindOfOrder      string             `bson:"kindOfOrder,omitempty"`
	KindOfOffer      string             `bson:"kindOfOffer,omitempty"`
	OrderType        string             `bson:"orderType,omitempty"`
	Commodity        primitive.ObjectID `bson:"commodity,omitempty"`
	AccountCode      string             `bson:"accountCode,omitempty"`
	Volume1Lot       string             `bson:"volume1Lot,omitempty"`
	Packing          string             `bson:"packing,omitempty"`
	DepositRate      int32              `bson:"depositRate,omitempty"`
	TransactionType  string             `bson:"transactionType,omitempty"`
	DeliveryTime     string             `bson:"deliveryTime,omitempty"`
	DeliveryLocation string             `bson:"deliveryLocation,omitempty"`
	Assessor         string             `bson:"assessor,omitempty"`
	Price            uint64             `bson:"price,omitempty"`
	CurrencyUnit     string             `bson:"currencyUnit,omitempty"`
	OrderVolume      string             `bson:"orderVolume,omitemty"`
	OrderValidity    int32              `bson:"orderValidity,omitempty"`
	Status           string             `bson:"status,omitempty"`
	CreatedAt        int64              `bson:"createdAt,omitempty"`
	CreatedBy        string             `bson:"createdBy,omitempty"`
	MatchedOrderID   primitive.ObjectID `bson:"matchedOrderId,omitempty"`
}
