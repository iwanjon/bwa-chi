package app

import (
	"bwastartupgochi/helper"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// ServerKey := os.Getenv("ServerKey")
	host := os.Getenv("host")
	port := os.Getenv("port")
	userdb := os.Getenv("userdb")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	newpoert, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Error convert port to int")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, newpoert, userdb, password, dbname)

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
