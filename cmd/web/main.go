package main

import (
	"log"
	"net/http"
	"sakura-internet-expt/config"
	"sakura-internet-expt/database"
	"sakura-internet-expt/handler"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)
	h := handler.NewHandler(db)
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
