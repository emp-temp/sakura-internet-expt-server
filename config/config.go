package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port   int    `env:"PORT" envDefault:"8080"`
	DBName string `env:"DB_NAME" envDefault:"sakuradb"`
	DBUser string `env:"DB_USER" envDefault:"sakura"`
	DBPass string `env:"DB_PASS" envDefault:"sakura-internet"`
	DBAddr string `env:"DB_ADDR" envDefault:"localhost:3306"`
}

func LoadConfig() *Config {
	if os.Getenv("Env") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &Config{
		Port:   port,
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBAddr: os.Getenv("DB_ADDR"),
	}
}
