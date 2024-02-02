package app

import (
	"bwastartupgochi/helper"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	userdb   = "postgres"
	password = "a"
	dbname   = "bwastartup"
)

func NewDB() *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, userdb, password, dbname)

	// db, err := sql.Open("mysql", "root:a@tcp(localhost:5555)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := sql.Open("postgres", psqlInfo)

	helper.PanicIfError(err, " error connection db")

	// defer db.Close()

	err = db.PingContext(ctx)

	helper.PanicIfError(err, " error in test ping database")

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// db.SetConnMaxIdleTime(10 * time.Minute)
	// defer func() {
	// 	fmt.Println("end the db")
	// 	db.Close()
	// 	fmt.Println("end the db r")
	// }()

	return db

}
