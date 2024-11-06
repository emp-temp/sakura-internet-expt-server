package main

import (
	"goweb/config"
	"goweb/database"
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)
	h := handler.NewHandler(db)
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
