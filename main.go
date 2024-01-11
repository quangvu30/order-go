package main

import (
	"fmt"

	"github.com/quangvu30/order-go/logger"
)

func main() {

	logging, err := logger.NewLogger()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logging.Info("Hello World")
}
