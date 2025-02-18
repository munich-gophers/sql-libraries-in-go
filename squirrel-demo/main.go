package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/davecgh/go-spew/spew"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	Age       int64
	Birthday  time.Time
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

func main() {
	// Set up Squirrel query builder
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Create a dynamic query using Squirrel
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, _, _ := psql.Select("*").From("users").ToSql()

	// Execute query
	rows, err := conn.Query(context.Background(), sql)
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
