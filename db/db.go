package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	dnsInput := "${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true"
	dns := os.ExpandEnv(dnsInput)

	var err error
	DB, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}


	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxIdleConns(3)
	DB.SetMaxOpenConns(3)

	if err = DB.Ping(); err != nil{
		log.Fatal(err)
	}
}
