package main

import (
	"fmt"

	"github.com/quangvu30/order-go/logger"
)

type Data struct {
	Name string
	Age  int
}

func main() {

	logging, err := logger.NewLogger()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logging.Info(Data{
		Name: "QUang",
		Age:  18,
	})
}
