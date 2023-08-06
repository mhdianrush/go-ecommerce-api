package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-ecommerce-api/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods(http.MethodGet)
}
