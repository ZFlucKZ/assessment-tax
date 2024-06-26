package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", Env.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
 
	log.Println("Connected to database")

	return db
}