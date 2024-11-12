package main

import (
	"database/sql"
	"log"
	"sakura-internet-expt/config"
	"sakura-internet-expt/database"
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("failed to close db")
		}
	}(db)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error migrating: %v", err)
	}
}
