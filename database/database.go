package database

import (
	"database/sql"
	"log"
	"sakura-internet-expt/config"
	"time"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	l, _ := time.LoadLocation("Asia/Tokyo")
	c := mysql.Config{
		DBName:    cfg.DBName,
		User:      cfg.DBUser,
		Passwd:    cfg.DBPass,
		Net:       "tcp",
		Addr:      cfg.DBAddr,
		ParseTime: true,
		Loc:       l,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("sql.Open err: %v", err)
	}
	return db
}
