package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type DB struct {
	Driver   string `env:"DB_DRIVER"`
	Database string `env:"DB_DATABASE"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Address  string `env:"DB_ADDRESS"`
	Port     string `env:"DB_PORT"`
	DSN      string
}

func CreateDBInfo() *DB {
	var db DB

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	if err := env.Parse(&db); err != nil {
		panic(err)
	}

	db.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db.User, db.Password, db.Address, db.Port, db.Database)

	return &db
}

func CreateTestDBInfo() *DB {
	db := DB{
		Driver:   "postgres",
		Database: "test",
		User:     "root",
		Password: "root",
		Address:  "localhost",
		Port:     "9999",
	}

	db.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db.User, db.Password, db.Address, db.Port, db.Database)

	return &db
}
