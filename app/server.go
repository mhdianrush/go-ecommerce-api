package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-ecommerce-api/database/seeders"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
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

type DatabaseConfig struct {
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabasePort     string
}

func (server *Server) InitializeDatabase(databaseConfig DatabaseConfig) {
	var err error
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		databaseConfig.DatabaseHost, databaseConfig.DatabaseUser, databaseConfig.DatabasePassword, databaseConfig.DatabaseName, databaseConfig.DatabasePort,
	)

	server.DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Printf("failed connected to database %s", err.Error())
	}

	for _, entity := range RegisterEntities() {
		err = server.DB.Debug().AutoMigrate(entity.Entity)
		if err != nil {
			logger.Printf("failed auto migrate database with gorm %s", err.Error())
		}
	}
	logger.Println("Success Migrate Database With GORM")
}

func (server *Server) Initialize(databaseConfig DatabaseConfig) {
	server.InitializeDatabase(databaseConfig)
	server.InitializeRoutes()
	seeders.DatabaseSeed(server.DB)
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
	var databaseConfig = DatabaseConfig{}

	err := godotenv.Load()
	if err != nil {
		logger.Printf("error loaded env file %s", err.Error())
	}

	appConfig.AppName = getEnv("APP_NAME", "failed load app name")
	appConfig.AppEnv = getEnv("APP_ENV", "failed load app env")
	appConfig.AppPort = getEnv("APP_PORT", "failed load app port")

	databaseConfig.DatabaseHost = getEnv("DATABASE_HOST", "failed load database host")
	databaseConfig.DatabaseUser = getEnv("DATABASE_USER", "failed load database user")
	databaseConfig.DatabasePassword = getEnv("DATABASE_PASSWORD", "failed load database password")
	databaseConfig.DatabaseName = getEnv("DATABASE_NAME", "failed load database name")
	databaseConfig.DatabasePort = getEnv("DATABASE_PORT", "failed load database port")

	server.Initialize(databaseConfig)
	server.Run(":" + appConfig.AppPort)
}
