package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"sqlc-demo/db" // sqlc generated DB code

	"github.com/davecgh/go-spew/spew"
	pgx "github.com/jackc/pgx/v5"
)

func main() {
	// Initialize sqlc-generated DB connection
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer conn.Close(ctx)

	// Example query using sqlc-generated code
	queries := db.New(conn)
	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		log.Fatalf("error querying all users %s\n", err)
	}

	fmt.Printf("Query Results (%d)\n", len(users))
	for _, entry := range users {
		spew.Dump(entry)
	}
}
