package winrate

import (
	"database/sql"
	"fmt"
	"log"

    _ "github.com/lib/pq" // driver to PostGreSQL
)

// Configuration to PostgreSQL driver
const (
	host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "your_password"
    dbname   = "your_database"
)

// Connection to database
func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password-%s dbname=%s ssqlmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error to connect database: %v", err)
	}

	if err = db.ping(); err != nil {
		log.Fatalf("Error to verify connection: %v", err)
	}

	log.Println("Connection established!")
	return db
}
