package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID        int
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	Age       int64
	Birthday  time.Time
}

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Example query
	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Fatalf("failed running Query %s\n", err)
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		log.Fatalf("collectrows caused an error: %v\n", err)
	}
	fmt.Printf("Query Results (%d)\n", len(users))
	for _, entry := range users {
		spew.Dump(entry)
	}
}
