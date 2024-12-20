package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Users struct {
	ID    int64
	Email string
}

func SqlConn() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "db",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func EmailAdd(users Users, db *sql.DB) {
	result, err := db.Exec("insert into usersemail email values ?", users.Email)
	if err != nil {
		fmt.Printf("EmailAdd: %v", err)
		return
	} else {
		fmt.Print(result)
	}
}
