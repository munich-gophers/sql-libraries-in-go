package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type User struct {
	ID         int
	Created_At *time.Time
	Deleted_At *time.Time
	Updated_At *time.Time
	Name       string
	Age        int64
	Birthday   time.Time
}

func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	// Example query
	var users []User
	err = db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatalf("couldn't fetch from database %s\n", err)
	}

	fmt.Printf("Query Results (%d)\n", len(users))
	for _, entry := range users {
		spew.Dump(entry)
	}
}
