package main

import (
	"log"
	"net/http"
	"practicegolangweb/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.ResgisterRoutes(mux)
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
