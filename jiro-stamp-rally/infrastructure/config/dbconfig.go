package dbconfig

import (
	"os"
	"strconv"
)

type DBconfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func NewDBConfig() *DBconfig {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &DBconfig{
		Host:     os.Getenv("HOST"),
		Port:     port,
		DBName:   os.Getenv("DBNAME"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
	}
}
