package main

import (
	"os"

	"github.com/mhdianrush/go-ecommerce-api/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Printf("failed created log file %s", err.Error())
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	app.Run()
}
