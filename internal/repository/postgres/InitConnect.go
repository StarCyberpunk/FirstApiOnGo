package postgres

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func CreateConnection() (*sql.DB,error) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	
	if err != nil {
		return nil,err
	}
	// check the connection
	err = db.Ping()

	if err != nil {
		return nil,err
	}
	return db,err
}
