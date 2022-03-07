package database


import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)


var DB *sql.DB
var err error

// Grab environment variables and connect to the database
func ConnectDB() {

	host := os.Getenv("HOST")
	port := 5432
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Printf("!@#")
		panic(err)
	}
}
