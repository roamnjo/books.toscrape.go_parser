package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func Connect() (*Storage, error) {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Open db error:", err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	rating VARCHAR(100) NOT NULL,
	price VARCHAR(100) NOT NULL,
	link VARCHAR(100) NOT NULL)
	`)
	defer stmt.Close()

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveBooks(title, rating, price, link string) (int64, error) {
	var id int64

	err := s.db.QueryRow(`INSERT INTO books(title, rating, price, link) 
	VALUES($1, $2, $3, $4) RETURNING id`,
		title, rating, price, link).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("SaveBooks: error:", err)
	}
	return id, nil
}
