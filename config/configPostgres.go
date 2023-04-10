package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func InitPostgres() {

	var err error
	DB, err = sql.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		panic(err)
	}
	//defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

}
