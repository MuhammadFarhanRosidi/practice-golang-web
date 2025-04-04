package routes

import (
	"net/http"
	"practicegolangweb/handlers"
)

func ResgisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/products", handlers.ProductsHandler)
	mux.HandleFunc("/articles", handlers.ArticlesHandler)
}
