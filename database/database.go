package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		name,
	)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Cannot connect to MySQL:", err)
	}

	fmt.Println("Connected to Cambria Bank MySQL database")
}
