package main

import (
	"goweb/config"
	"goweb/database"
	"log"
)

var query = `
CREATE TABLE IF NOT EXISTS cds_data (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	start_time DATETIME NOT NULL,
	end_time DATETIME NOT NULL
);
`

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)
	defer db.Close()
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error migrating: %v", err)
	}
}
