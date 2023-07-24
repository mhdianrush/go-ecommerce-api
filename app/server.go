package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var logger = logrus.New()

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	logger.Printf("Server Running on Port %s", address)
	err := http.ListenAndServe(address, server.Router)
	if err != nil {
		logger.Printf("failed connected to server %s", err.Error())
	}
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	appConfig.AppName = "ECommerceWeb"
	appConfig.AppEnv = "development"
	appConfig.AppPort = "8080"

	server.Initialize()
	server.Run(":" + appConfig.AppPort)
}
