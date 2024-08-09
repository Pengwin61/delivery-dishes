package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClientRoach struct {
	Con *gorm.DB
}

var (
	host     = "localhost"
	port     = "5432"
	user     = "delivery"
	password = "123456"
	dbname   = "delivery"
)

func NewClientCockroach() *ClientRoach {
	strDSN := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(strDSN), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &ClientRoach{Con: db}
}
