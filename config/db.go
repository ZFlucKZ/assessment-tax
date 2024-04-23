package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDB() (*DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
 
	log.Println("Connected to database")

	return &DB{db}, nil
}