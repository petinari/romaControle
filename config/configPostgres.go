package config

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

var DB *pgx.Conn

func InitPostgres() {

	var err error
	DB, err = pgx.Connect(context.Background(), os.Getenv("POSTGRES"))
	if err != nil {
		panic(err)
	}

	//defer DB.Close()
	err = DB.Ping(context.Background())
	if err != nil {
		panic(err)
	}

}
