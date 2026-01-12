package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database not responding:", err)
	}

	DB = db
	CreateExpensesTable()
}

func CreateExpensesTable() {
	query := `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			amount NUMERIC(10, 2) NOT NULL,
			category VARCHAR(100) NOT NULL,
			description TEXT,
			date VARCHAR(50),
			user_id INTEGER NOT NULL
		)
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create expenses table:", err)
	}

	log.Println("Expenses table created or already exists")
}
