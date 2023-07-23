package app

import (
	"net/http"

	"github.com/mhdianrush/go-ecommerce-api/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
}
