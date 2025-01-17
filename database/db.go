package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// type dbClient struct {
// 	client *sql.DB
// }

// func (c *dbClient) Query(query string) (*sql.Rows, error) {
// 	return c.client.Query(query)
// }

// func (c *dbClient) Exec(query string) (sql.Result, error) {
// 	return c.client.Exec(query)
// }

var pgClient *sql.DB

func InitDB() error {
	//connStr := "host=localhost port=5432 user=harsh password=harsh dbname=mydb sslmode=disable"

	s := os.Getenv("pg_connection")
	if s == "" {
		log.Fatalf("Environment variable 'pg_connection' is not set")
		return fmt.Errorf("environment variable 'pg_connection' is not set")
	}

	db, err := sql.Open("postgres", s)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return err
	}
	fmt.Print("connected to DB...")
	pgClient = db

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return nil
}
