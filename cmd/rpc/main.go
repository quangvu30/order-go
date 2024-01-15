package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/quangvu30/order-go/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderTransferClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateOrder(ctx, &pb.Order{
		OrderID:          "23234324dsfsdf",
		KindOfOrder:      "new",
		KindOfOffer:      "offer",
		OrderType:        "sell",
		Commodity:        "SVR7",
		AccountCode:      "024I000045",
		Volume1Lot:       "657c0603806720d5ae738351",
		DepositRate:      10,
		TransactionType:  "forward",
		DeliveryTime:     32,
		DeliveryLocation: "TP HCM",
		Price:            40000,
		OrderVolume:      "12",
		Packing:          "pallet",
		Assessor:         "buyer",
		OrderValidity:    24,
		Status:           "fom-order",
		CreatedAt:        12234235345,
		CreatedBy:        "quangvd@mxv.vn",
		MatchedOrderID:   "sdfsdf3325"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.OrderID)
}
