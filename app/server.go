package app

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		logger.Printf("error loaded env file %s", err.Error())
	}

	appConfig.AppName = getEnv("APP_NAME", "failed load app name")
	appConfig.AppEnv = getEnv("APP_ENV", "failed load app env")
	appConfig.AppPort = getEnv("APP_PORT", "failed load app port")

	server.Initialize()
	server.Run(":" + appConfig.AppPort)
}
