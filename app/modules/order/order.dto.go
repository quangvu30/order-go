package order

import "github.com/quangvu30/order-go/config"

type ReqCreateOrder struct {
	KindOfOrder      string `json:"kindOfOrder"`
	KindOfOffer      string `json:"kindOfOffer"`
	OrderType        string `json:"orderType"`
	Commodity        string `json:"commodity"`
	AccountCode      string `json:"accountCode"`
	Volume1Lot       string `json:"volume1Lot"`
	Packing          string `json:"packing"`
	DepositRate      int32  `json:"depositRate"`
	TransactionType  string `json:"transactionType"`
	DeliveryTime     uint64 `json:"deliveryTime"`
	DeliveryLocation string `json:"deliveryLocation"`
	Assessor         string `json:"assessor"`
	Price            uint64 `json:"price"`
	CurrencyUnit     string `json:"currencyUnit"`
	OrderVolume      string `json:"orderVolum"`
	OrderValidity    int32  `json:"orderValidity"`
}

type ResCreateOrder struct {
	config.Response
}
