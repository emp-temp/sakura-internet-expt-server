package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"sakura-internet-expt/config"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	c := mysql.Config{
		DBName:    cfg.DBName,
		User:      cfg.DBUser,
		Passwd:    cfg.DBPass,
		Net:       "tcp",
		Addr:      cfg.DBAddr,
		ParseTime: true,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("sql.Open err: %v", err)
	}
	return db
}
